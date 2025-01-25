package service

import (
	"context"
	"github.com/saleh-ghazimoradi/GoInn/internal/repository"
	"github.com/saleh-ghazimoradi/GoInn/internal/service/service_models"
)

type UserService interface {
	GetUserById(ctx context.Context, id string) (*service_models.User, error)
	GetUsers(ctx context.Context) ([]*service_models.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func (u *userService) GetUserById(ctx context.Context, id string) (*service_models.User, error) {
	return u.userRepository.GetUserById(ctx, id)
}

func (u *userService) GetUsers(ctx context.Context) ([]*service_models.User, error) {
	return u.userRepository.GetUsers(ctx)
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}
