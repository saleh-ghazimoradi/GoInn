package repository

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/GoInn/internal/service/service_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id primitive.ObjectID) (*service_models.User, error)
	GetUsers(ctx context.Context) ([]*service_models.User, error)
	CreateUser(ctx context.Context, user *service_models.User) (*service_models.User, error)
	UpdateUser(ctx context.Context, user *service_models.User) (*service_models.User, error)
	DeleteUser(ctx context.Context, id primitive.ObjectID) error
}

type userRepository struct {
	collection *mongo.Collection
}

func (u *userRepository) CreateUser(ctx context.Context, user *service_models.User) (*service_models.User, error) {
	res, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, errors.New("error inserting user")
	}

	user.Id = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (u *userRepository) GetUserById(ctx context.Context, id primitive.ObjectID) (*service_models.User, error) {
	var user service_models.User
	if err := u.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (u *userRepository) GetUsers(ctx context.Context) ([]*service_models.User, error) {
	cursor, err := u.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []*service_models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, errors.New("failed to fetch users")
	}

	return users, nil
}

func (u *userRepository) UpdateUser(ctx context.Context, user *service_models.User) (*service_models.User, error) {
	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": user}
	_, err := u.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, errors.New("failed to update user")
	}

	return user, nil
}

func (u *userRepository) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	_, err := u.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}
