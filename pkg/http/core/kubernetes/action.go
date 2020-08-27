package kubernetes

import (
	"github.com/billiford/go-clouddriver/pkg/kubernetes"
	"github.com/billiford/go-clouddriver/pkg/sql"
	"github.com/gin-gonic/gin"
)

//go:generate counterfeiter . Action
type Action interface {
	Run() error
}

type ActionConfig struct {
	KubeClient  kubernetes.Client
	SQLClient   sql.Client
	ID          string
	Application string
	Operation   Operation
}

//go:generate counterfeiter . ActionHandler
type ActionHandler interface {
	NewDeployManifestAction(ActionConfig) Action
	NewRollingRestartAction(ActionConfig) Action
	NewRollbackAction(ActionConfig) Action
	NewScaleManifestAction(ActionConfig) Action
}

const ActionHandlerInstanceKey = `KubernetesActionHandler`

func NewActionHandler() ActionHandler {
	return &actionHandler{}
}

type actionHandler struct{}

func ActionHandlerInstance(c *gin.Context) ActionHandler {
	return c.MustGet(ActionHandlerInstanceKey).(ActionHandler)
}
