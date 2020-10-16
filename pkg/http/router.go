package http

import (
	"github.com/billiford/go-clouddriver/pkg/http/core"
	v1 "github.com/billiford/go-clouddriver/pkg/http/v1"
	"github.com/gin-gonic/gin"
)

// Define the API.
func Initialize(r *gin.Engine) {
	// API endpoints without a version will go under "core".
	api := r.Group("")
	{
		api.GET("/health", core.OK)

		// Force cache refresh.
		api.POST("/cache/kubernetes/manifest", core.OK)

		// Credentials API controller.
		api.GET("/credentials", core.ListCredentials)
		api.GET("/credentials/:account", core.GetAccountCredentials)

		// Applications API controller.
		//
		// https://github.com/spinnaker/clouddriver/blob/master/clouddriver-web/src/main/groovy/com/netflix/spinnaker/clouddriver/controllers/ApplicationsController.groovy#L38
		// @PreAuthorize("#restricted ? @fiatPermissionEvaluator.storeWholePermission() : true")
		// @PostFilter("#restricted ? hasPermission(filterObject.name, 'APPLICATION', 'READ') : true")
		api.GET("/applications", core.ListApplications)

		// https://github.com/spinnaker/clouddriver/blob/master/clouddriver-web/src/main/groovy/com/netflix/spinnaker/clouddriver/controllers/ServerGroupManagerController.java#L39
		// @PreAuthorize("hasPermission(#application, 'APPLICATION', 'READ')")
		// @PostFilter("hasPermission(filterObject.account, 'ACCOUNT', 'READ')")
		api.GET("/applications/:application/serverGroupManagers", core.ListServerGroupManagers)

		// https://github.com/spinnaker/clouddriver/blob/master/clouddriver-web/src/main/groovy/com/netflix/spinnaker/clouddriver/controllers/ServerGroupController.groovy#L172
		// @PreAuthorize("hasPermission(#application, 'APPLICATION', 'READ')")
		// @PostAuthorize("@authorizationSupport.filterForAccounts(returnObject)")
		api.GET("/applications/:application/serverGroups", core.ListServerGroups)

		// https://github.com/spinnaker/clouddriver/blob/master/clouddriver-web/src/main/groovy/com/netflix/spinnaker/clouddriver/controllers/ServerGroupController.groovy#L75
		// @PreAuthorize("hasPermission(#account, 'ACCOUNT', 'READ')")
		// @PostAuthorize("hasPermission(returnObject?.moniker?.app, 'APPLICATION', 'READ')")
		api.GET("/applications/:application/serverGroups/:account/:location/:name", core.GetServerGroup)

		// https: //github.com/spinnaker/clouddriver/blob/master/clouddriver-web/src/main/groovy/com/netflix/spinnaker/clouddriver/controllers/LoadBalancerController.groovy#L42
		// @PreAuthorize("hasPermission(#application, 'APPLICATION', 'READ')")
		// @PostAuthorize("@authorizationSupport.filterForAccounts(returnObject)")
		api.GET("/applications/:application/loadBalancers", core.ListLoadBalancers)

		// https://github.com/spinnaker/clouddriver/blob/master/clouddriver-web/src/main/groovy/com/netflix/spinnaker/clouddriver/controllers/ClusterController.groovy#L44
		// @PreAuthorize("@fiatPermissionEvaluator.storeWholePermission() and hasPermission(#application, 'APPLICATION', 'READ')")
		// @PostAuthorize("@authorizationSupport.filterForAccounts(returnObject)")
		api.GET("/applications/:application/clusters", core.ListClusters)

		// https://github.com/spinnaker/clouddriver/blob/master/clouddriver-web/src/main/groovy/com/netflix/spinnaker/clouddriver/controllers/JobController.groovy#L35
		// @PreAuthorize("hasPermission(#application, 'APPLICATION', 'READ') and hasPermission(#account, 'ACCOUNT', 'READ')")
		// @ApiOperation(value = "Collect a JobStatus", notes = "Collects the output of the job.")
		api.GET("/applications/:application/jobs/:account/:location/:name", core.GetJob)

		// Create a kubernetes operation - deploy/delete/scale manifest.
		api.POST("/kubernetes/ops", core.CreateKubernetesOperation)

		// Manifests API controller.
		api.GET("/manifests/:account/:location/:kind", core.GetManifest)
		api.GET("/manifests/:account/:location/:kind/cluster/:application/:cluster/dynamic/:target", core.GetManifestByTarget)

		// Get results for a task triggered in CreateKubernetesOperation.
		api.GET("/task/:id", core.GetTask)

		// Generic search endpoint.
		r.GET("/search", core.Search)

		// Not implemented.
		api.GET("/securityGroups", core.ListSecurityGroups)

		// Artifacts API controller.
		api.GET("/artifacts/credentials", core.ListArtifactCredentials)
		api.GET("/artifacts/account/:accountName/names", core.ListHelmArtifactAccountNames)
		api.GET("/artifacts/account/:accountName/versions", core.ListHelmArtifactAccountVersions)
		api.PUT("/artifacts/fetch/", core.GetArtifact)

		// Features.
		api.GET("/features/stages", core.ListStages)
	}

	// New endpoint.
	api = r.Group("/v1")
	{
		// Providers endpoint for kubernetes.
		api.POST("/kubernetes/providers", v1.CreateKubernetesProvider)
	}
}
