package api

import (
	"net/http"

	_ "github.com/rkpundhir90/task-management/docs"

	"github.com/gin-gonic/gin"
	"github.com/rkpundhir90/task-management/task-management/api/handler"
	"github.com/rkpundhir90/task-management/task-management/app"
	"github.com/rkpundhir90/task-management/task-management/config"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Task Management API
// @version 1.0
// @description API for managing tasks (CRUD operations).
// @host localhost:8080
// @BasePath /v1

// NewRouter sets up the routes and returns the Gin engine
func NewRouter(a app.Interface, envConfig *config.EnvironmentConfiguration) *gin.Engine {

	var taskInstanceService = a.TaskInstanceService()
	var userInstanceService = a.UserInstanceService()

	router := gin.Default()
	router.Use(gin.Recovery())

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "Method not available for this route"})
	})

	v1 := router.Group("/v1")

	handler.NewRegisterTaskHandlers(envConfig, v1, taskInstanceService)
	handler.NewRegisterUserHandlers(envConfig, v1, userInstanceService)

	_ = router.Run(":" + envConfig.ServerPort)

	return router
}
