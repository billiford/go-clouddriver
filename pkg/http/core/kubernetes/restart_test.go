package kubernetes_test

import (
	"errors"

	"github.com/billiford/go-clouddriver/pkg/kubernetes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Restart", func() {
	BeforeEach(func() {
		setup()
		action = actionHandler.NewRollingRestartAction(actionConfig)
	})

	JustBeforeEach(func() {
		err = action.Run()
	})

	When("getting the provider returns an error", func() {
		BeforeEach(func() {
			fakeSQLClient.GetKubernetesProviderReturns(kubernetes.Provider{}, errors.New("error getting provider"))
		})

		It("returns an error", func() {
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("error getting provider"))
		})
	})

	When("there is an error decoding the CA data for the kubernetes provider", func() {
		BeforeEach(func() {
			fakeSQLClient.GetKubernetesProviderReturns(kubernetes.Provider{CAData: "{}{}{}{}"}, nil)
		})

		It("returns an error", func() {
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("illegal base64 data at input byte 0"))
		})
	})

	When("creating the kube client returns an error", func() {
		BeforeEach(func() {
			fakeKubeController.NewClientReturns(nil, errors.New("bad config"))
		})

		It("returns an error", func() {
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("bad config"))
		})
	})

	When("getting the manifest returns an error", func() {
		BeforeEach(func() {
			fakeKubeClient.GetReturns(nil, errors.New("error getting manifest"))
		})

		It("returns an error", func() {
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("error getting manifest"))
		})
	})

	When("applying the manifest returns an error", func() {
		BeforeEach(func() {
			fakeKubeClient.ApplyReturns(kubernetes.Metadata{}, errors.New("error applying manifest"))
		})

		It("returns an error", func() {
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("error applying manifest"))
		})
	})

	When("it succeeds", func() {
		It("succeeds", func() {
			Expect(err).To(BeNil())
		})
	})
})
