package repository

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/GoInn/internal/service/service_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelRepository interface {
	InsertHotel(ctx context.Context, hotel *service_models.Hotel) (*service_models.Hotel, error)
}

type hotelRepository struct {
	collection *mongo.Collection
}

func (h *hotelRepository) InsertHotel(ctx context.Context, hotel *service_models.Hotel) (*service_models.Hotel, error) {
	res, err := h.collection.InsertOne(ctx, hotel)
	if err != nil {
		return nil, errors.New("error inserting hotel")
	}
	hotel.Id = res.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func NewHotelRepository(db *mongo.Database) HotelRepository {
	return &hotelRepository{
		collection: db.Collection("hotels"),
	}
}
