package repository

import (
	"context"

	"github.com/rkpundhir90/task-management/task-management/config"
	"github.com/rkpundhir90/task-management/task-management/model"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type UserInstanceRepository interface {
	GetUserByEmailId(ctx context.Context, emailId string) (model.User, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, user model.CreateUserRequest) error
}

func NewUserInstanceRepository(envConfig *config.EnvironmentConfiguration, db *gorm.DB) UserInstanceRepository {
	return repository{envConfig, db}
}

func (r repository) GetUserByEmailId(ctx context.Context, emailId string) (model.User, error) {
	var user model.User
	result := r.db.WithContext(ctx).Table("users").First(&user, "email = ?", emailId)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("error getting product by key")
		return model.User{}, result.Error
	}
	return user, nil
}

func (r repository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	result := r.db.WithContext(ctx).Table("users").Find(&users)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("error getting product by key")
		return []model.User{}, result.Error
	}
	return users, nil
}

func (r repository) CreateUser(ctx context.Context, user model.CreateUserRequest) error {
	result := r.db.WithContext(ctx).Table("users").Create(&user)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("error getting user by key")
		return result.Error
	}
	return nil
}
