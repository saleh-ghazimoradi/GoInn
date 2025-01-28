package service

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/GoInn/helper"
	"github.com/saleh-ghazimoradi/GoInn/internal/dto"
	"github.com/saleh-ghazimoradi/GoInn/internal/repository"
	"github.com/saleh-ghazimoradi/GoInn/internal/service/service_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	GetUserById(ctx context.Context, id string) (*service_models.User, error)
	GetUsers(ctx context.Context) ([]*service_models.User, error)
	CreateUser(ctx context.Context, user *dto.User) (*service_models.User, error)
	UpdateUser(ctx context.Context, id string, input *dto.UpdateUser) (*service_models.User, error)
	DeleteUser(ctx context.Context, id string) error
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
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}
	return u.userRepository.GetUserById(ctx, oid)
}

func (u *userService) GetUsers(ctx context.Context) ([]*service_models.User, error) {
	return u.userRepository.GetUsers(ctx)
}

func (u *userService) UpdateUser(ctx context.Context, id string, input *dto.UpdateUser) (*service_models.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	existingUser, err := u.userRepository.GetUserById(ctx, oid)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if input.FirstName != nil {
		existingUser.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		existingUser.LastName = *input.LastName
	}

	updatedUser, err := u.userRepository.UpdateUser(ctx, existingUser)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (u *userService) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID")
	}
	return u.userRepository.DeleteUser(ctx, oid)
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}
