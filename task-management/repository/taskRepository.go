package repository

import (
	"context"

	"github.com/rkpundhir90/task-management/task-management/config"
	"github.com/rkpundhir90/task-management/task-management/model"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type repository struct {
	envConfig *config.EnvironmentConfiguration
	db        *gorm.DB
}

type TaskInstanceRepository interface {
	GetTaskById(ctx context.Context, id int) (model.Task, error)
	GetAllTasks(ctx context.Context) ([]model.Task, error)
	DeleteTaskById(ctx context.Context, id int) error
	SaveTask(ctx context.Context, task model.CreateTaskRequest) error
	UpdateTaskById(ctx context.Context, task model.UpdateTaskRequest) error
}

func NewTaskInstanceRepository(envConfig *config.EnvironmentConfiguration, db *gorm.DB) TaskInstanceRepository {
	return repository{envConfig, db}
}

func (r repository) GetTaskById(ctx context.Context, id int) (model.Task, error) {
	var task model.Task
	result := r.db.WithContext(ctx).Table("tasks").First(&task, "id = ?", id)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("error getting product by key")
		return model.Task{}, result.Error
	}
	return task, nil
}

func (r repository) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	var task []model.Task
	result := r.db.WithContext(ctx).Table("tasks").Find(&task)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("error getting product by key")
		return []model.Task{}, result.Error
	}
	return task, nil
}

func (r repository) DeleteTaskById(ctx context.Context, id int) error {
	result := r.db.WithContext(ctx).Table("tasks").Delete(&model.Task{}, "id = ?", id)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("error getting product by key")
		return result.Error
	}
	return nil
}

func (r repository) SaveTask(ctx context.Context, task model.CreateTaskRequest) error {
	result := r.db.WithContext(ctx).Table("tasks").Create(&task)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("error getting product by key")
		return result.Error
	}
	return nil
}

func (r repository) UpdateTaskById(ctx context.Context, task model.UpdateTaskRequest) error {

	result := r.db.WithContext(ctx).Table("tasks").Save(&task)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("error getting product by key")
		return result.Error
	}
	return nil
}
