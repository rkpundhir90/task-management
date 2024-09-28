package service

import (
	"context"

	"github.com/rkpundhir90/task-management/task-management/config"
	"github.com/rkpundhir90/task-management/task-management/model"
	"github.com/rkpundhir90/task-management/task-management/repository"
)

type TaskInstanceService interface {
	GetTaskById(ctx context.Context, taskId int) (model.TaskInfo, error)
}

type taskInstanceService struct {
	envConfig  *config.EnvironmentConfiguration
	repository repository.TaskInstanceRepository
}

func NewTaskInstanceService(
	envConfig *config.EnvironmentConfiguration,
	repository repository.TaskInstanceRepository,
) TaskInstanceService {
	return taskInstanceService{
		envConfig,
		repository,
	}
}

func (s taskInstanceService) GetTaskById(ctx context.Context, taskId int) (model.TaskInfo, error) {
	return model.TaskInfo{TaskId: taskId}, nil
}
