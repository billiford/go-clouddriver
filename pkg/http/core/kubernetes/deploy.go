package kubernetes

import (
	"encoding/base64"

	"github.com/billiford/go-clouddriver/pkg/kubernetes"
	"github.com/billiford/go-clouddriver/pkg/sql"
	"github.com/google/uuid"
	"k8s.io/client-go/rest"
)

func (ah *actionHandler) NewDeployManifestAction(ac ActionConfig) Action {
	return &deployManfest{
		sc: ac.SQLClient,
		kc: ac.KubeClient,
		id: ac.ID,
		dm: ac.Operation.DeployManifest,
	}
}

type deployManfest struct {
	sc sql.Client
	kc kubernetes.Client
	id string
	dm *DeployManifestRequest
}

func (d *deployManfest) Run() error {
	provider, err := d.sc.GetKubernetesProvider(d.dm.Account)
	if err != nil {
		return err
	}

	cd, err := base64.StdEncoding.DecodeString(provider.CAData)
	if err != nil {
		return err
	}

	config := &rest.Config{
		Host:        provider.Host,
		BearerToken: provider.BearerToken,
		TLSClientConfig: rest.TLSClientConfig{
			CAData: cd,
		},
	}

	d.kc.WithConfig(config)

	for _, manifest := range d.dm.Manifests {
		u, err := kubernetes.ToUnstructured(manifest)
		if err != nil {
			return err
		}

		err = kubernetes.AddSpinnakerAnnotations(u, d.dm.Moniker.App)
		if err != nil {
			return err
		}

		err = kubernetes.AddSpinnakerLabels(u, d.dm.Moniker.App)
		if err != nil {
			return err
		}

		meta, err := d.kc.Apply(u)
		if err != nil {
			return err
		}

		kr := kubernetes.Resource{
			AccountName:  d.dm.Account,
			ID:           uuid.New().String(),
			TaskID:       d.id,
			APIGroup:     meta.Group,
			Name:         meta.Name,
			Namespace:    meta.Namespace,
			Resource:     meta.Resource,
			Version:      meta.Version,
			Kind:         meta.Kind,
			SpinnakerApp: d.dm.Moniker.App,
		}

		err = d.sc.CreateKubernetesResource(kr)
		if err != nil {
			return err
		}
	}

	return nil
}
