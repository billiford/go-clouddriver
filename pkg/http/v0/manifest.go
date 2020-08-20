package v0

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	clouddriver "github.com/billiford/go-clouddriver/pkg"
	ops "github.com/billiford/go-clouddriver/pkg/http/v0/kubernetes"
	"github.com/billiford/go-clouddriver/pkg/kubernetes"
	"github.com/billiford/go-clouddriver/pkg/sql"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/rest"
)

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
		BearerToken: provider.BearerToken,
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

	app := "unknown"
	labels := result.GetLabels()
	if _, ok := labels["app.kubernetes.io/name"]; ok {
		app = labels["app.kubernetes.io/name"]
	}

	kmr := ops.ManifestResponse{
		Account:  account,
		Events:   []interface{}{},
		Location: namespace,
		Manifest: result.Object,
		Metrics:  []interface{}{},
		Moniker: ops.ManifestResponseMoniker{
			App:     app,
			Cluster: fmt.Sprintf("%s %s", kind, name),
		},
		Name: fmt.Sprintf("%s %s", kind, name),
		// The 'default' status of a kubernetes resource.
		Status:   kubernetes.GetStatus(kind, result.Object),
		Warnings: []interface{}{},
	}

	c.JSON(http.StatusOK, kmr)
}
