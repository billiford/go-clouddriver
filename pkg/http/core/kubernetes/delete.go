package kubernetes

import (
	"fmt"
	"strings"

	"github.com/billiford/go-clouddriver/pkg/arcade"
	"github.com/billiford/go-clouddriver/pkg/kubernetes"
	"github.com/billiford/go-clouddriver/pkg/sql"
)

// Example requests.
//
// Choosing a dynamic manifest to delete.
// [
//   {
//     "deleteManifest": {
//       "app": "smoketests",
//       "mode": "dynamic",
//       "cluster": "deployment delete-me",
//       "criteria": "newest",
//       "kind": "deployment",
//       "cloudProvider": "kubernetes",
//       "manifestName": "deployment delete-me",
//       "options": {
//         "cascading": true
//       },
//       "location": "smoketest",
//       "account": "gke_github-replication-sandbox_us-central1_sandbox-us-central1-agent_smoketest-dev"
//     }
//   }
// ]
//
// Choosing manifests to delete by label.
// selectors can be:
// ANY, EQUALS, NOT_EQUALS, CONTAINS, NOT_CONTAINS, EXISTS, NOT_EXISTS
// [
//   {
//     "deleteManifest": {
//       "app": "smoketests",
//       "mode": "label",
//       "cloudProvider": "kubernetes",
//       "labelSelectors": {
//         "selectors": [
//           {
//             "kind": "EQUALS",
//             "values": [
//               "fdsaasdffew"
//             ],
//             "key": "asdf"
//           }
//         ]
//       },
//       "options": {
//         "cascading": true
//       },
//       "kinds": [
//         "deployment"
//       ],
//       "location": "smoketest",
//       "account": "gke_github-replication-sandbox_us-central1_sandbox-us-central1-agent_smoketest-dev"
//     }
//   }
// ]
func (ah *actionHandler) NewDeleteManifestAction(ac ActionConfig) Action {
	return &deleteManfest{
		ac: ac.ArcadeClient,
		sc: ac.SQLClient,
		kc: ac.KubeController,
		id: ac.ID,
		dm: ac.Operation.DeleteManifest,
	}
}

type deleteManfest struct {
	ac arcade.Client
	sc sql.Client
	kc kubernetes.Controller
	id string
	dm *DeleteManifestRequest
}

// TODO waiting on https://github.com/billiford/go-clouddriver/pull/10 which has client
// functions needed to dynamically delete by cluster.
func (d *deleteManfest) Run() error {
	// provider, err := d.sc.GetKubernetesProvider(d.dm.Account)
	// if err != nil {
	// 	return err
	// }
	//
	// cd, err := base64.StdEncoding.DecodeString(provider.CAData)
	// if err != nil {
	// 	return err
	// }
	//
	// token, err := d.ac.Token()
	// if err != nil {
	// 	return err
	// }

	// config := &rest.Config{
	// 	Host:        provider.Host,
	// 	BearerToken: token,
	// 	TLSClientConfig: rest.TLSClientConfig{
	// 		CAData: cd,
	// 	},
	// }

	// client, err := d.kc.NewClient(config)
	// if err != nil {
	// 	return err
	// }

	switch strings.ToLower(d.dm.Mode) {
	case "dynamic":
	default:
		return fmt.Errorf("requested to delete manifest %s using mode %s which is not implemented", d.dm.ManifestName, d.dm.Mode)
	}

	return nil
}
