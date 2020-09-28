package kubernetes_test

import (
	. "github.com/billiford/go-clouddriver/pkg/kubernetes"
	"github.com/billiford/go-clouddriver/pkg/kubernetes/manifest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/apps/v1"
	//corev1 "k8s.io/api/core/v1"
	//v1 "k8s.io/api/apps/v1beta1"
)

var _ = Describe("Statefulset", func() {
	var (
		ss  StatefulSet
		//err error
	)

	BeforeEach(func() {
		s := map[string]interface{}{}
		ss = NewStatefulSet(s)
	})

	/*Describe("#GetStatefulSetSpec", func() {
		BeforeEach(func() {
			_ = ss.GetStatefulSetSpec()
		})

		It("succeeds", func() {
		})
	})

	Describe("#GetStatefulSetStatus", func() {
		BeforeEach(func() {
			_ = ss.GetStatefulSetStatus()
		})

		It("succeeds", func() {
		})
	})*/

	Describe("#Object", func() {
		var s *v1.StatefulSet

		BeforeEach(func() {
			s = ss.Object()
		})

		When("it succeeds", func() {
			It("succeeds", func() {
				Expect(s).ToNot(BeNil())
			})
		})
	})


        Describe("#Status", func() {
                var s manifest.Status

                BeforeEach(func() {
                        replicas := int32(4)
                        ss.SetReplicas(&replicas)
                })

                JustBeforeEach(func() {
                        s = ss.Status()
                })

		When("there are more desired replicas than ready", func() {
			BeforeEach(func() {
				o := ss.Object()
				o.Status.ReadyReplicas = int32(4)
			})

			It("returns status unstable", func() {
				Expect(s.Stable.State).To(BeFalse())
				Expect(s.Stable.Message).To(Equal("Waiting for current replicas in stateful set to match expected replicas"))
			})
		})
	})
})
