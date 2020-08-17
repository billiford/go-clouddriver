package v0_test

import (
	// . "github.com/billiford/go-clouddriver/pkg/http/v0"

	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ok", func() {
	BeforeEach(func() {
		setup()
		uri = svr.URL + "/health"
		createRequest(http.MethodGet)
	})

	AfterEach(func() {
		svr.Close()
	})

	JustBeforeEach(func() {
		doRequest()
	})

	It("returns OK", func() {
		Expect(res.StatusCode).To(Equal(http.StatusOK))
	})
})
