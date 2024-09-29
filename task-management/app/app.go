package app

import (
	"github.com/rkpundhir90/task-management/task-management/config"
	"github.com/rkpundhir90/task-management/task-management/repository"
	"github.com/rkpundhir90/task-management/task-management/service"
	"gorm.io/gorm"
)

var _ Interface = (*App)(nil)

type Interface interface {
	TaskInstanceService() service.TaskInstanceService
	UserInstanceService() service.UserInstanceService

	TaskInstanceRepository() repository.TaskInstanceRepository
	UserInstanceRepository() repository.UserInstanceRepository
}

type App struct {
	taskInstanceService    service.TaskInstanceService
	userInstanceService    service.UserInstanceService
	taskInstanceRepository repository.TaskInstanceRepository
	userInstanceRepository repository.UserInstanceRepository
}

func (a *App) TaskInstanceRepository() repository.TaskInstanceRepository {
	return a.taskInstanceRepository
}

func (a *App) TaskInstanceService() service.TaskInstanceService {
	return a.taskInstanceService
}

func (a *App) UserInstanceService() service.UserInstanceService {
	return a.userInstanceService
}

func (a *App) UserInstanceRepository() repository.UserInstanceRepository {
	return a.userInstanceRepository
}

func Build(db *gorm.DB, envConfig *config.EnvironmentConfiguration) *App {
	taskInstanceRepository := repository.NewTaskInstanceRepository(envConfig, db)
	userInstanceRepository := repository.NewUserInstanceRepository(envConfig, db)

	taskInstanceService := service.NewTaskInstanceService(envConfig, taskInstanceRepository)
	userInstanceService := service.NewUserInstanceService(envConfig, userInstanceRepository)

	app := &App{
		taskInstanceRepository: taskInstanceRepository,
		taskInstanceService:    taskInstanceService,
		userInstanceRepository: userInstanceRepository,
		userInstanceService:    userInstanceService,
	}

	return app
}
