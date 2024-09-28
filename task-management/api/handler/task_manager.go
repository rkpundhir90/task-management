package handler

import (
	"net/http"

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
	r.GET("/task/:taskId/tasks", res.GetTasksByTaskId)

}

func (r TaskManagementResource) GetTasksByTaskId(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	ctx.JSON(http.StatusOK, taskId)
}
