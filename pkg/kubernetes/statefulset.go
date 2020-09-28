package kubernetes

import (
	"encoding/json"
	"strings"
	"github.com/billiford/go-clouddriver/pkg/kubernetes/manifest"
	v1 "k8s.io/api/apps/v1"
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

	updStrategyType := x.Spec.UpdateStrategy.Type
	if strings.EqualFold(updStrategyType, "OnDelete") {
		return s
	}

	if x.ObjectMeta.Generation != x.Status.ObservedGeneration {
		s.Stable.State = false
		s.Stable.Message = "Waiting for status generation to match updated object generation"
		return s
	}

	desired := int32(0)
	current := x.Status.CurrentReplicas
	ready := x.Status.ReadyReplicas
	existing := x.Status.Replicas

	if x.Spec.Replicas != nil {
		desired = *x.Spec.Replicas
	}

	if desired > existing {
		s.Stable.State = false
		s.Stable.Message = "Waiting for at least the desired replica count to be met"

		return s
	}

	if desired > ready {
		s.Stable.State = false
		s.Stable.Message = "Waiting for all updated replicas to be ready"

		return s
	}

	if desired > current {
		s.Stable.State = false
		s.Stable.Message = "Waiting for all updated replicas to be scheduled"

		return s
	}

	updateRev := x.Status.UpdateRevision
	currentRev := x.Status.CurrentRevision

	if currentRev != updateRev {
		s.Stable.State = false
		s.Stable.Message = "Waiting for the updated revision to match the current revision"
		return s	
	}

	return s
}
