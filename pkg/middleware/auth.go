package middleware

import (
	"net/http"

	clouddriver "github.com/billiford/go-clouddriver/pkg"
	"github.com/billiford/go-clouddriver/pkg/fiat"
	"github.com/gin-gonic/gin"
)

const (
	user = `X-Spinnaker-User`
	app  = `X-Spinnaker-Application`
)

//authApplication takes a list of permissions
//authAccount takes a list of accounts

func AuthApplication(permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.GetHeader(user)
		app := c.GetHeader(app)

		if user == "" || app == "" {
			c.Next()
			return
		}

		fiatClient := fiat.Instance(c)
		authResp, err := fiatClient.Authorize(user)
		if err != nil {
			clouddriver.WriteError(c, http.StatusUnauthorized, err)
			return
		}
		applicationsAuth := authResp.Applications
		for _, auth := range applicationsAuth {
			if auth.Name == app {
				for _, p := range permissions {
					found := find(auth.Authorizations, p)
					if !found {
						clouddriver.WriteError(c, http.StatusForbidden, nil)
						return
					}
				}
			}
		}
		c.Next()
	}
}

func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
