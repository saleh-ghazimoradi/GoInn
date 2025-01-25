package repository

import (
	"context"
	"github.com/saleh-ghazimoradi/GoInn/internal/service/service_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id string) (*service_models.User, error)
	GetUsers(ctx context.Context) ([]*service_models.User, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func (u *userRepository) GetUserById(ctx context.Context, id string) (*service_models.User, error) {
	var user service_models.User
	if err := u.collection.FindOne(ctx, bson.M{"_id": primitive.ObjectIDFromHex(id)}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) GetUsers(ctx context.Context) ([]*service_models.User, error) {
	cursor, err := u.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []*service_models.User
	if err = cursor.Decode(&users); err != nil {
		return nil, err
	}
	
	return users, nil
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}
