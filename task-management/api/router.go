package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rkpundhir90/task-management/task-management/api/handler"
	"github.com/rkpundhir90/task-management/task-management/app"
	"github.com/rkpundhir90/task-management/task-management/config"
)

// NewRouter sets up the routes and returns the Gin engine
func NewRouter(a app.Interface, envConfig *config.EnvironmentConfiguration) *gin.Engine {

	var taskInstanceService = a.TaskInstanceService()

	router := gin.Default()
	router.Use(gin.Recovery())

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "Method not available for this route"})
	})

	v1 := router.Group("/v1")

	handler.NewRegisterTaskHandlers(envConfig, v1, taskInstanceService)

	return router
}
