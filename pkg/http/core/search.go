package core

import (
	"errors"
	"net/http"
	"strconv"

	clouddriver "github.com/billiford/go-clouddriver/pkg"
	"github.com/gin-gonic/gin"
)

type SearchResponse []Page

type Page struct {
	PageNumber int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
	// Platform     string       `json:"platform"`
	Query        string       `json:"query"`
	Results      []PageResult `json:"results"`
	TotalMatches int          `json:"totalMatches"`
}

type PageResult struct {
	Account        string `json:"account"`
	Group          string `json:"group"`
	KubernetesKind string `json:"kubernetesKind"`
	Name           string `json:"name"`
	Namespace      string `json:"namespace"`
	Provider       string `json:"provider"`
	Region         string `json:"region"`
	Type           string `json:"type"`
	Application    string `json:"application,omitempty"`
	Cluster        string `json:"cluster,omitempty"`
}

func Search(c *gin.Context) {
	sr := SearchResponse{}
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	namespace := c.Query("q")
	// The "type" query param is the kubernetes kind.
	kind := c.Query("type")
	accounts := c.GetHeader("X-Spinnaker-Accounts")

	if kind == "" || namespace == "" {
		clouddriver.WriteError(c, http.StatusBadRequest, errors.New("must provide query params 'q' and 'type'"))
		return
	}

	p := Page{}
	results := []PageResult{}
	for _, account := range accounts {

	}

	c.JSON(http.StatusOK, []string{})
}
