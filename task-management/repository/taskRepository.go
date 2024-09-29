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
