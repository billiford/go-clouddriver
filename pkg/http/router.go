package http

import (
	v0 "github.com/billiford/go-clouddriver/pkg/http/v0"
	v1 "github.com/billiford/go-clouddriver/pkg/http/v1"
	"github.com/gin-gonic/gin"
)

// Define the API.
func Initialize(r *gin.Engine) {
	// API endpoints without a version will go under "v0".
	api := r.Group("")
	{
		api.GET("/health", v0.OK)
		// Force cache refresh.
		api.POST("/cache/kubernetes/manifest", v0.OK)
		api.GET("/credentials", v0.ListCredentials)
		api.GET("/applications/:application/serverGroupManagers", v0.ListServerGroupManagers)

		// Trigger a kubernetes deployment.
		r.POST("/kubernetes/ops", v0.CreateKubernetesDeployment)

		// Monitor deploy.
		r.GET("/manifests/:account/:location/:name", v0.GetManifest)

		r.GET("/task/:id", v0.GetTask)
	}

	api = r.Group("/v1")
	{
		api.POST("/kubernetes/providers", v1.CreateKubernetesProvider)
	}
}
