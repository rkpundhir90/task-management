package service

import (
	"context"

	"github.com/rkpundhir90/task-management/task-management/config"
	"github.com/rkpundhir90/task-management/task-management/model"
	"github.com/rkpundhir90/task-management/task-management/repository"
)

type UserInstanceService interface {
	GetUserByEmailId(ctx context.Context, emailId string) (model.User, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, user model.CreateUserRequest) error
}

type userInstanceService struct {
	envConfig  *config.EnvironmentConfiguration
	repository repository.UserInstanceRepository
}

func (u userInstanceService) CreateUser(ctx context.Context, user model.CreateUserRequest) error {
	err := u.repository.CreateUser(ctx, user)
	return err
}

func (u userInstanceService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	users, err := u.repository.GetAllUsers(ctx)
	return users, err
}

func (u userInstanceService) GetUserByEmailId(ctx context.Context, emailId string) (model.User, error) {
	user, err := u.repository.GetUserByEmailId(ctx, emailId)
	return user, err
}

func NewUserInstanceService(
	envConfig *config.EnvironmentConfiguration,
	repository repository.UserInstanceRepository,
) UserInstanceService {
	return userInstanceService{
		envConfig,
		repository,
	}
}
