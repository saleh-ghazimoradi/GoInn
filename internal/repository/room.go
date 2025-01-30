package repository

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/GoInn/internal/service/service_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomRepository interface {
	InsertRoom(ctx context.Context, room *service_models.Room) (*service_models.Room, error)
}

type roomRepository struct {
	collection *mongo.Collection
}

func (r *roomRepository) InsertRoom(ctx context.Context, room *service_models.Room) (*service_models.Room, error) {
	res, err := r.collection.InsertOne(ctx, room)
	if err != nil {
		return nil, errors.New("error creating room")
	}

	room.Id = res.InsertedID.(primitive.ObjectID)
	return room, nil
}

func NewRoomRepository(db *mongo.Database) RoomRepository {
	return &roomRepository{
		collection: db.Collection("rooms"),
	}
}
