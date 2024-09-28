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

	TaskInstanceRepository() repository.TaskInstanceRepository
}

type App struct {
	taskInstanceService    service.TaskInstanceService
	taskInstanceRepository repository.TaskInstanceRepository
}

func (a *App) TaskInstanceRepository() repository.TaskInstanceRepository {
	return a.taskInstanceRepository
}

func (a *App) TaskInstanceService() service.TaskInstanceService {
	return a.taskInstanceService
}

func Build(db *gorm.DB, envConfig *config.EnvironmentConfiguration) *App {
	taskInstanceRepository := repository.NewTaskInstanceRepository(envConfig, db)

	taskInstanceService := service.NewTaskInstanceService(envConfig, taskInstanceRepository)

	app := &App{
		taskInstanceRepository: taskInstanceRepository,
		taskInstanceService:    taskInstanceService,
	}

	return app
}
