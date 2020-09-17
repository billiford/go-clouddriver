package main

import (
	"log"
	"os"

	"github.com/billiford/go-clouddriver/pkg/arcade"
	"github.com/billiford/go-clouddriver/pkg/helm"
	kube "github.com/billiford/go-clouddriver/pkg/http/core/kubernetes"
	"github.com/billiford/go-clouddriver/pkg/kubernetes"
	"github.com/billiford/go-clouddriver/pkg/server"
	"github.com/billiford/go-clouddriver/pkg/sql"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/mcuadros/go-gin-prometheus"
)

var (
	r = gin.New()
)

func main() {
	r.Run(":7002")
}

func init() {
	// Setup metrics.
	p := ginprometheus.NewPrometheus("gin")
	p.MetricsPath = "/metrics"
	p.Use(r)

	gin.ForceConsoleColor()
	// Ignore logging of certain endpoints.
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/health"}}))
	r.Use(gin.Recovery())

	sqlConfig := sql.Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Name:     os.Getenv("DB_NAME"),
	}

	db, err := sql.Connect(sql.Connection(sqlConfig))
	if err != nil {
		log.Fatal(err.Error())
	}

	sqlClient := sql.NewClient(db)
	kubeController := kubernetes.NewController()
	helmClient := helm.NewClient("https://kubernetes-charts.storage.googleapis.com")
	arcadeClient := arcade.NewDefaultClient()

	arcadeAPIKey := os.Getenv("ARCADE_API_KEY")
	if arcadeAPIKey == "" {
		log.Println("WARNING: ARCADE_API_KEY not set")
	}

	arcadeClient.WithAPIKey(arcadeAPIKey)

	c := &server.Config{
		ArcadeClient:      arcadeClient,
		SQLClient:         sqlClient,
		KubeController:    kubeController,
		KubeActionHandler: kube.NewActionHandler(),
		HelmClient:        helmClient,
	}
	if os.Getenv("VERBOSE_REQUEST_LOGGING") == "true" {
		c.VerboseRequestLogging = true
	}

	server.Setup(r, c)
}
