package v0

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	clouddriver "github.com/billiford/go-clouddriver/pkg"
	"github.com/billiford/go-clouddriver/pkg/kubernetes"
	"github.com/billiford/go-clouddriver/pkg/kubernetes/manifest"
	"github.com/billiford/go-clouddriver/pkg/kubernetes/pod"
	"github.com/billiford/go-clouddriver/pkg/kubernetes/replicaset"
	"github.com/billiford/go-clouddriver/pkg/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"k8s.io/client-go/rest"
)

func CreateKubernetesDeployment(c *gin.Context) {
	sc := sql.Instance(c)
	kc := kubernetes.Instance(c)
	taskID := uuid.New().String()
	kor := OpsRequest{}

	err := c.ShouldBindJSON(&kor)
	if err != nil {
		clouddriver.WriteError(c, http.StatusBadRequest, err)
		return
	}

	// TODO this is hacky - need to figure out how to handle providers.
	accountName := kor[0].DeployManifest.Account

	provider, err := sc.GetKubernetesProvider(kor[0].DeployManifest.Account)
	if err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	cd, err := base64.StdEncoding.DecodeString(provider.CAData)
	if err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	config := &rest.Config{
		Host:        provider.Host,
		BearerToken: os.Getenv("BEARER_TOKEN"),
		TLSClientConfig: rest.TLSClientConfig{
			CAData: cd,
		},
	}

	if err = kc.WithConfig(config); err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	for _, req := range kor {
		for _, manifest := range req.DeployManifest.Manifests {
			b, err := json.Marshal(manifest)
			if err != nil {
				clouddriver.WriteError(c, http.StatusBadRequest, err)
				return
			}

			_, meta, err := kc.Apply(b)
			if err != nil {
				clouddriver.WriteError(c, http.StatusInternalServerError, err)
				return
			}

			kr := kubernetes.Resource{
				AccountName: accountName,
				ID:          uuid.New().String(),
				TaskID:      taskID,
				APIGroup:    meta.Group,
				Name:        meta.Name,
				Namespace:   meta.Namespace,
				Resource:    meta.Resource,
				Version:     meta.Version,
				Kind:        meta.Kind,
			}

			err = sc.CreateKubernetesResource(kr)
			if err != nil {
				clouddriver.WriteError(c, http.StatusInternalServerError, err)
				return
			}
		}
	}

	or := OpsResponse{
		ID:          taskID,
		ResourceURI: "/task/" + taskID,
	}
	c.JSON(http.StatusOK, or)
}

func GetManifest(c *gin.Context) {
	sc := sql.Instance(c)
	kc := kubernetes.Instance(c)
	account := c.Param("account")
	namespace := c.Param("location")
	n := c.Param("name")
	a := strings.Split(n, " ")
	kind := a[0]
	name := a[1]

	provider, err := sc.GetKubernetesProvider(account)
	if err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	cd, err := base64.StdEncoding.DecodeString(provider.CAData)
	if err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	config := &rest.Config{
		Host:        provider.Host,
		BearerToken: os.Getenv("BEARER_TOKEN"),
		TLSClientConfig: rest.TLSClientConfig{
			CAData: cd,
		},
	}

	if err = kc.WithConfig(config); err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	result, err := kc.Get(kind, name, namespace)
	if err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	kmr := ManifestResponse{
		Account:  account,
		Events:   nil,
		Location: namespace,
		Manifest: result.Object,
		Metrics:  nil,
		Moniker: struct {
			App     string "json:\"app\""
			Cluster string "json:\"cluster\""
		}{
			App:     "TODO",
			Cluster: "TODO",
		},
		Name: name,
		// The 'default' status of a kubernetes resource.
		Status:   manifest.DefaultStatus,
		Warnings: nil,
	}

	// status https://github.com/spinnaker/clouddriver/tree/master/clouddriver-kubernetes/src/main/java/com/netflix/spinnaker/clouddriver/kubernetes/op/handler
	// pod status check https://github.com/spinnaker/clouddriver/blob/master/clouddriver-kubernetes/src/main/java/com/netflix/spinnaker/clouddriver/kubernetes/op/handler/KubernetesPodHandler.java
	switch strings.ToLower(kind) {
	case "pod":
		kmr.Status = pod.Status(result.Object)
	case "replicaset":
		kmr.Status = replicaset.Status(result.Object)
	default:
		kmr.Status = manifest.DefaultStatus
	}

	c.JSON(http.StatusOK, kmr)
}
