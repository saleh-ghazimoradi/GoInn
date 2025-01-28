package service

import (
	"context"
	"github.com/saleh-ghazimoradi/GoInn/helper"
	"github.com/saleh-ghazimoradi/GoInn/internal/dto"
	"github.com/saleh-ghazimoradi/GoInn/internal/repository"
	"github.com/saleh-ghazimoradi/GoInn/internal/service/service_models"
)

type UserService interface {
	GetUserById(ctx context.Context, id string) (*service_models.User, error)
	GetUsers(ctx context.Context) ([]*service_models.User, error)
	CreateUser(ctx context.Context, user *dto.User) (*service_models.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func (u *userService) CreateUser(ctx context.Context, user *dto.User) (*service_models.User, error) {
	hashedPassword, err := helper.CreateHashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	createdUser, err := u.userRepository.CreateUser(ctx, &service_models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		return nil, err
	}

	return createdUser, nil
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
