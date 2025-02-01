package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/saleh-ghazimoradi/GoInn/internal/service/service_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomRepository interface {
	InsertRoom(ctx context.Context, room *service_models.Room) (*service_models.Room, error)
	GetRooms(ctx context.Context) ([]*service_models.Room, error)
	GetRoomsByHotelId(ctx context.Context, hotelId primitive.ObjectID) ([]*service_models.Room, error)
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

func (r *roomRepository) GetRooms(ctx context.Context) ([]*service_models.Room, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New("error getting rooms")
	}
	var rooms []*service_models.Room
	if err = cursor.All(ctx, &rooms); err != nil {
		return nil, errors.New("error getting rooms")
	}
	return rooms, nil
}

func (r *roomRepository) GetRoomsByHotelId(ctx context.Context, hotelId primitive.ObjectID) ([]*service_models.Room, error) {
	// Use the ObjectID directly in the query
	filter := bson.M{"hotelId": hotelId}

	fmt.Println("Querying rooms with filter:", filter) // Debug logging

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error querying rooms: %v", err)
	}

	var rooms []*service_models.Room
	if err = cursor.All(ctx, &rooms); err != nil {
		return nil, fmt.Errorf("error decoding rooms: %v", err)
	}

	return rooms, nil
}

func NewRoomRepository(db *mongo.Database) RoomRepository {
	return &roomRepository{
		collection: db.Collection("rooms"),
	}
}
