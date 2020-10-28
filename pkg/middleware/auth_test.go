package middleware_test

import (
	. "github.com/billiford/go-clouddriver/pkg/middleware"
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("Auth", func() {
	Context("#AuthApplication", func() {
		BeforeEach(func() {
			AuthApplication("READ")
		})
	})
})
