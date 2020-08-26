package core

import (
	"log"
	"net/http"

	clouddriver "github.com/billiford/go-clouddriver/pkg"
	"github.com/billiford/go-clouddriver/pkg/http/core/kubernetes"
	kube "github.com/billiford/go-clouddriver/pkg/kubernetes"
	"github.com/billiford/go-clouddriver/pkg/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// The main function that starts a kubernetes operation.
//
// Kubernetes operations are things like deploy/delete manifest or perform
// a rolling restart. Spinnaker sends *all* of these types of events to the
// same endpoint (/kubernetes/ops), so we have to unmarshal and check which
// kind of operation we are performing.
//
// The actual actions have been moved to the kubernetes subfolder to make
// this function a bit more readable.
func CreateKubernetesOperation(c *gin.Context) {
	// All operations are bound to a task ID and stored in the database.
	taskID := uuid.New().String()
	ko := kubernetes.Operations{}

	err := c.ShouldBindJSON(&ko)
	if err != nil {
		clouddriver.WriteError(c, http.StatusBadRequest, err)
		return
	}

	// Spinnaker likes to send an 'extra' POST request to /kubernetes/ops -
	// I have not figured out what these requests are yet. I'll need to unmarshal
	// into a map[string]interface{} in order to read all the fields being sent.
	//
	// For now, I return status OK for this task - so far so good!
	if len(ko) == 0 {
		or := kubernetes.OperationsResponse{
			ID:          taskID,
			ResourceURI: "/task/" + taskID,
		}
		c.JSON(http.StatusOK, or)
		return
	}

	// Loop through each request in the kubernetes operations and perform
	// each requested action.
	for _, req := range ko {
		if req.DeployManifest != nil {
			action := kubernetes.NewDeployManifestAction(
				sql.Instance(c),
				kube.Instance(c),
				taskID,
				*req.DeployManifest,
			)
			err = action.Run()
			if err != nil {
				clouddriver.WriteError(c, http.StatusInternalServerError, err)
				return
			}
		}

		if req.ScaleManifest != nil {
			action := kubernetes.NewScaleManifestAction(
				sql.Instance(c),
				kube.Instance(c),
				taskID,
				*req.ScaleManifest,
			)
			err = action.Run()
			if err != nil {
				clouddriver.WriteError(c, http.StatusInternalServerError, err)
				return
			}
		}

		if req.CleanupArtifacts != nil {
			log.Println("got request to cleanup artifacts - unimplemented")
		}

		if req.RollingRestartManifest != nil {
			action := kubernetes.NewRollingRestartAction(
				sql.Instance(c),
				kube.Instance(c),
				taskID,
				*req.RollingRestartManifest,
			)
			err = action.Run()
			if err != nil {
				clouddriver.WriteError(c, http.StatusInternalServerError, err)
				return
			}
		}

		if req.UndoRolloutManifest != nil {
			action := kubernetes.NewRollbackAction(
				sql.Instance(c),
				kube.Instance(c),
				taskID,
				c.GetHeader("X-Spinnaker-Application"),
				*req.UndoRolloutManifest,
			)
			err = action.Run()
			if err != nil {
				clouddriver.WriteError(c, http.StatusInternalServerError, err)
				return
			}
		}
	}

	or := kubernetes.OperationsResponse{
		ID:          taskID,
		ResourceURI: "/task/" + taskID,
	}
	c.JSON(http.StatusOK, or)
}
