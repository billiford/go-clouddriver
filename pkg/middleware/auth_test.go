package middleware_test

import (
	"github.com/billiford/go-clouddriver/pkg/middleware"
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("Auth", func() {
	var (
		permissions []string
	)
	Context("#AuthApplication", func() {
		BeforeEach(func() {
			permissions = []string{"READ"}
			auth := middleware.NewPermissions(permissions)
			auth.AuthApplication()
		})
	})
})
