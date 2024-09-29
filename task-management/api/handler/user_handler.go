package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rkpundhir90/task-management/task-management/config"
	"github.com/rkpundhir90/task-management/task-management/model"
	"github.com/rkpundhir90/task-management/task-management/service"
)

type UserManagementResource struct {
	envConfig   *config.EnvironmentConfiguration
	userService service.UserInstanceService
}

func NewRegisterUserHandlers(envConfig *config.EnvironmentConfiguration, r *gin.RouterGroup, userService service.UserInstanceService) {
	res := UserManagementResource{envConfig, userService}

	// Register all routes with the handler functions
	r.GET("/user", res.GetAllUsers)
	r.GET("/user/:emailId", res.GetUserByEmailId)
	r.POST("/user", res.CreateUser)
}

// GetAllUsers retrieves all the users
// @Summary Retrieve all the users
// @Description Retrieve all the users
// @Tags users
// @Produce json
// @Success 200 {array} []model.User
// @Router /user [get]
func (r UserManagementResource) GetAllUsers(ctx *gin.Context) {

	users, err := r.userService.GetAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// GetUserByEmailId get user by emailId
// @Summary Get User by email id
// @Description Get User by email id
// @Tags users
// @Produce json
// @Param emailId path string true "Email ID"
// @Success 200 {object} model.User
// @Router /user/{emailId} [get]
func (r UserManagementResource) GetUserByEmailId(ctx *gin.Context) {
	userIdParam := ctx.Param("emailId")
	user, err := r.userService.GetUserByEmailId(ctx, userIdParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// CreateUser create a user
// @Summary Create a user
// @Description Create a user
// @Tags users
// @Produce json
// @Param   user  body  model.CreateUserRequest  true  "User to be created"
// @Success 201
// @Router /user [post]
func (r UserManagementResource) CreateUser(ctx *gin.Context) {

	var user model.CreateUserRequest

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := r.userService.CreateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
