package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rkpundhir90/task-management/task-management/config"
	"github.com/rkpundhir90/task-management/task-management/service"
)

type TaskManagementResource struct {
	envConfig   *config.EnvironmentConfiguration
	taskService service.TaskInstanceService
}

func NewRegisterTaskHandlers(envConfig *config.EnvironmentConfiguration, r *gin.RouterGroup, taskService service.TaskInstanceService) {
	res := TaskManagementResource{envConfig, taskService}

	// Register all routes with the handler functions
	r.GET("/tasks/:taskId", res.GetTasksByTaskId)

}

// GetTasksByTaskId retrieves the task by Id
// @Summary Retrieve task by Id
// @Description Retrieve task by Id
// @Tags tasks
// @Produce json
// @Param taskId path int true "Task ID"
// @Success 200 {array} model.Task
// @Router /tasks/{taskId} [get]
func (r TaskManagementResource) GetTasksByTaskId(ctx *gin.Context) {
	taskIdParam := ctx.Param("taskId")
	taskId, err := strconv.Atoi(taskIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	task, err := r.taskService.GetTaskById(ctx, taskId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}
