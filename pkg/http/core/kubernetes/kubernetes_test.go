package kubernetes_test

import (
	"github.com/billiford/go-clouddriver/pkg/arcade/arcadefakes"
	. "github.com/billiford/go-clouddriver/pkg/http/core/kubernetes"
	"github.com/billiford/go-clouddriver/pkg/kubernetes"
	"github.com/billiford/go-clouddriver/pkg/kubernetes/kubernetesfakes"
	"github.com/billiford/go-clouddriver/pkg/sql/sqlfakes"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	err                error
	fakeArcadeClient   *arcadefakes.FakeClient
	fakeSQLClient      *sqlfakes.FakeClient
	fakeKubeClient     *kubernetesfakes.FakeClient
	fakeKubeController *kubernetesfakes.FakeController
	actionConfig       ActionConfig
	actionHandler      ActionHandler
	action             Action
)

func setup() {
	// Setup fakes.
	fakeArcadeClient = &arcadefakes.FakeClient{}

	fakeSQLClient = &sqlfakes.FakeClient{}
	fakeSQLClient.GetKubernetesProviderReturns(kubernetes.Provider{
		Name:   "test-account",
		Host:   "http://localhost",
		CAData: "",
	}, nil)
	fakeSQLClient.ListKubernetesResourcesByTaskIDReturns([]kubernetes.Resource{
		{
			AccountName: "test-account-name",
		},
	}, nil)

	ul := &unstructured.UnstructuredList{
		Items: []unstructured.Unstructured{
			{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							kubernetes.AnnotationSpinnakerArtifactName: "test-deployment",
							kubernetes.AnnotationSpinnakerArtifactType: "kubernetes/deployment",
							"deployment.kubernetes.io/revision":        "100",
						},
					},
				},
			},
		},
	}
	fakeKubeClient = &kubernetesfakes.FakeClient{}
	fakeKubeClient.GetReturns(&unstructured.Unstructured{Object: map[string]interface{}{}}, nil)
	fakeKubeClient.ListByGVRReturns(ul, nil)

	fakeKubeController = &kubernetesfakes.FakeController{}
	fakeKubeController.NewClientReturns(fakeKubeClient, nil)

	actionHandler = NewActionHandler()
	actionConfig = newActionConfig()
}

func newActionConfig() ActionConfig {
	return ActionConfig{
		ArcadeClient:   fakeArcadeClient,
		KubeController: fakeKubeController,
		SQLClient:      fakeSQLClient,
		ID:             "test-id",
		Application:    "test-application",
		Operation: Operation{
			DeployManifest: &DeployManifestRequest{
				Manifests: []map[string]interface{}{
					{
						"kind":       "Pod",
						"apiVersion": "v1",
					},
				},
			},
			ScaleManifest: &ScaleManifestRequest{
				Replicas:     "16",
				ManifestName: "deployment test-deployment",
			},
			CleanupArtifacts: &CleanupArtifactsRequest{},
			DeleteManifest: &DeleteManifestRequest{
				ManifestName: "deployment test-deployment",
			},
			UndoRolloutManifest: &UndoRolloutManifestRequest{
				ManifestName: "deployment test-deployment",
				Revision:     "100",
			},
			RollingRestartManifest: &RollingRestartManifestRequest{
				ManifestName: "deployment test-deployment",
			},
			PatchManifest: &PatchManifestRequest{
				ManifestName: "deployment test-deployment",
			},
		},
	}
}
