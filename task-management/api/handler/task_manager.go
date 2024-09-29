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
	r.GET("/tasks/:taskId/tasks", res.GetTasksByTaskId)

}

// GetTasks retrieves all tasks or filters by status, user, or due date
// @Summary Retrieve tasks
// @Description Retrieve all tasks or filter tasks by status, user, or due date.
// @Tags tasks
// @Produce json
// @Param status query string false "Status"
// @Param assigned_to query string false "Assigned user"
// @Param due_date query string false "Due date"
// @Success 200 {array} model.TaskInfo
// @Router /tasks [get]
func (r TaskManagementResource) GetTasksByTaskId(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	ctx.JSON(http.StatusOK, taskId)
}
