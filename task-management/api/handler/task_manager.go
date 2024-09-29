package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rkpundhir90/task-management/task-management/config"
	"github.com/rkpundhir90/task-management/task-management/model"
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
	r.GET("/tasks", res.GetAllTasks)
	r.POST("/tasks", res.CreateTask)
	r.DELETE("/tasks/:taskId", res.DeleteTaskById)
	r.PUT("/tasks", res.UpdateTaskById)

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

// GetAllTasks retrieves all the tasks
// @Summary Retrieve All Tasks
// @Description Retrieve All Tasks
// @Tags tasks
// @Produce json
// @Success 200 {array} []model.Task
// @Router /tasks [get]
func (r TaskManagementResource) GetAllTasks(ctx *gin.Context) {
	task, err := r.taskService.GetAllTasks(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

// DeleteTaskById delete task by Id
// @Summary Delete Task by Id
// @Description Delete Task by Id
// @Tags tasks
// @Produce json
// @Param taskId path int true "Task ID"
// @Success 204
// @Router /tasks/{taskId} [delete]
func (r TaskManagementResource) DeleteTaskById(ctx *gin.Context) {
	taskIdParam := ctx.Param("taskId")
	taskId, err := strconv.Atoi(taskIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	err = r.taskService.DeleteTaskById(ctx, taskId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// CreateTask create a task
// @Summary Create a Task
// @Description Create a Task
// @Tags tasks
// @Produce json
// @Param   task  body  model.CreateTaskRequest  true  "Task to be created"
// @Success 201
// @Router /tasks [post]
func (r TaskManagementResource) CreateTask(ctx *gin.Context) {

	var task model.CreateTaskRequest

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := r.taskService.CreateTask(ctx, task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// UpdateTaskById update task
// @Summary Update Task
// @Description Update Task
// @Tags tasks
// @Produce json
// @Param   task  body  model.UpdateTaskRequest  true  "Task to be updated"
// @Success 204
// @Router /tasks [put]
func (r TaskManagementResource) UpdateTaskById(ctx *gin.Context) {
	var task model.UpdateTaskRequest
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := r.taskService.UpdateTaskById(ctx, task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
