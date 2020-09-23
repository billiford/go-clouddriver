package kubernetes

import (
	"encoding/json"

	"github.com/billiford/go-clouddriver/pkg/kubernetes/manifest"
	//v1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/apps/v1beta1"
)

type StatefulSet interface {
	Object() *v1.StatefulSet
	SetReplicas(*int32)
	Status() manifest.Status
}

type statefulSet struct {
	ss *v1.StatefulSet
}

func NewStatefulSet(m map[string]interface{}) StatefulSet {
	s := &v1.StatefulSet{}
	b, _ := json.Marshal(m)
	_ = json.Unmarshal(b, &s)

	return &statefulSet{ss: s}
}

func (ss *statefulSet) Object() *v1.StatefulSet {
	return ss.ss
}

func (ss *statefulSet) SetReplicas(replicas *int32) {
	ss.ss.Spec.Replicas = replicas
}

func (ss *statefulSet) Status() manifest.Status {
	s := manifest.DefaultStatus
	x := ss.ss

	desired := int32(0)
	current := x.Status.CurrentReplicas
	ready := x.Status.ReadyReplicas

	if x.Spec.Replicas != nil {
		desired = *x.Spec.Replicas
	}

	if desired > ready {
		s.Stable.State = false
		s.Stable.Message = "Waiting for all replicas in stateful set to be ready"

		return s
	}

	if desired > current {
		s.Stable.State = false
		s.Stable.Message = "Waiting for current replicas in stateful set to match expected replicas"

		return s
	}

	//if x.ObjectMeta.Generation != x.Status.ObservedGeneration {
		//s.Stable.State = false
		//s.Stable.Message = "Waiting for statefulset spec update to be observed"

		//return s
	//}

	return s
}
