package kubernetes

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/billiford/go-clouddriver/pkg/kubernetes"
	"github.com/billiford/go-clouddriver/pkg/kubernetes/deployment"
	"github.com/billiford/go-clouddriver/pkg/sql"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/rest"
)

type ScaleManifestRequest struct {
	Replicas      string `json:"replicas"`
	ManifestName  string `json:"manifestName"`
	CloudProvider string `json:"cloudProvider"`
	Location      string `json:"location"`
	User          string `json:"user"`
	Account       string `json:"account"`
}

func ScaleManifest(c *gin.Context, sm ScaleManifestRequest) error {
	sc := sql.Instance(c)
	kc := kubernetes.Instance(c)

	provider, err := sc.GetKubernetesProvider(sm.Account)
	if err != nil {
		return err
	}

	cd, err := base64.StdEncoding.DecodeString(provider.CAData)
	if err != nil {
		return err
	}

	config := &rest.Config{
		Host:        provider.Host,
		BearerToken: os.Getenv("BEARER_TOKEN"),
		TLSClientConfig: rest.TLSClientConfig{
			CAData: cd,
		},
	}

	if err = kc.WithConfig(config); err != nil {
		return err
	}

	a := strings.Split(sm.ManifestName, " ")
	kind := a[0]
	name := a[1]

	unstructuredObj, err := kc.Get(kind, name, sm.Location)
	if err != nil {
		return err
	}

	switch strings.ToLower(kind) {
	case "deployment":
		d := deployment.New(unstructuredObj.Object)
		replicas, _ := strconv.Atoi(sm.Replicas)
		desiredReplicas := int32(replicas)

		d.Spec.Replicas = &desiredReplicas

		b, err := json.Marshal(d)
		if err != nil {
			return err
		}

		_, err = kc.Patch(b)
		if err != nil {
			return err
		}
	}

	return nil
}
