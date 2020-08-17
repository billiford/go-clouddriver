package middleware

import (
	"github.com/billiford/go-clouddriver/pkg/kubernetes"
	"github.com/billiford/go-clouddriver/pkg/sql"
	"github.com/gin-gonic/gin"
)

func SetSQLClient(r sql.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(sql.ClientInstanceKey, r)
		c.Next()
	}
}

func SetKubeClient(k kubernetes.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(kubernetes.ClientInstanceKey, k)
		c.Next()
	}
}
