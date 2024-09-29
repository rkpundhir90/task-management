package service

import (
	"context"

	"github.com/rkpundhir90/task-management/task-management/config"
	"github.com/rkpundhir90/task-management/task-management/model"
	"github.com/rkpundhir90/task-management/task-management/repository"
)

type TaskInstanceService interface {
	GetTaskById(ctx context.Context, taskId int) (model.Task, error)
	GetAllTasks(ctx context.Context) ([]model.Task, error)
	DeleteTaskById(ctx context.Context, taskId int) error
	CreateTask(ctx context.Context, task model.CreateTaskRequest) error
	UpdateTaskById(ctx context.Context, task model.UpdateTaskRequest) error
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

func (s taskInstanceService) GetTaskById(ctx context.Context, taskId int) (model.Task, error) {
	task, err := s.repository.GetTaskById(ctx, taskId)
	return task, err
}

func (s taskInstanceService) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	task, err := s.repository.GetAllTasks(ctx)
	return task, err
}

func (s taskInstanceService) DeleteTaskById(ctx context.Context, id int) error {
	err := s.repository.DeleteTaskById(ctx, id)
	return err
}

func (s taskInstanceService) CreateTask(ctx context.Context, task model.CreateTaskRequest) error {
	err := s.repository.SaveTask(ctx, task)
	return err
}

func (s taskInstanceService) UpdateTaskById(ctx context.Context, task model.UpdateTaskRequest) error {
	err := s.repository.UpdateTaskById(ctx, task)
	return err
}
