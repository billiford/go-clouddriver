package middleware

import (
	"net/http"

	clouddriver "github.com/billiford/go-clouddriver/pkg"
	"github.com/billiford/go-clouddriver/pkg/fiat"
	"github.com/gin-gonic/gin"
)

//authApplication takes a list of permissions
//authAccount
//go:generate counterfeiter . Auth
type Auth interface {
	AuthApplication() gin.HandlerFunc
}

type permissions struct {
	list []string
}

func NewPermissions(list []string) permissions {
	return permissions{list: list}
}

func (p *permissions) AuthApplication() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.GetHeader("user")
		app := c.GetHeader("application")

		if len(user) == 0 || len(app) == 0 {
			c.Next()
			return
		}

		fiatClient := fiat.NewDefaultClient()
		authResp, err := fiatClient.Authorize(user)
		if err != nil {
			clouddriver.WriteError(c, http.StatusUnauthorized, err)
			return
		}
		applicationsAuth := authResp.Applications
		for _, auth := range applicationsAuth {
			if auth.Name == app {
				for _, permission := range p.list {
					_, found := Find(auth.Authorizations, permission)
					if !found {
						clouddriver.WriteError(c, http.StatusForbidden, err)
						return
					}
				}
			}
		}
		c.Next()
	}
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
