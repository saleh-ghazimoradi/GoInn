package repository

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/GoInn/internal/service/service_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelRepository interface {
	InsertHotel(ctx context.Context, hotel *service_models.Hotel) (*service_models.Hotel, error)
	GetHotelById(ctx context.Context, id primitive.ObjectID) (*service_models.Hotel, error)
	GetHotels(ctx context.Context) ([]*service_models.Hotel, error)
	UpdateHotel(ctx context.Context, hotel *service_models.Hotel) (*service_models.Hotel, error)
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

func (h *hotelRepository) GetHotelById(ctx context.Context, id primitive.ObjectID) (*service_models.Hotel, error) {
	var hotel service_models.Hotel
	if err := h.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&hotel); err != nil {
		return nil, errors.New("hotel not found")
	}
	return &hotel, nil
}

func (h *hotelRepository) GetHotels(ctx context.Context) ([]*service_models.Hotel, error) {
	cursor, err := h.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New("error getting hotels")
	}

	var hotels []*service_models.Hotel
	if err = cursor.All(ctx, &hotels); err != nil {
		return nil, errors.New("failed to fetch hotels")
	}

	return hotels, nil
}

func (h *hotelRepository) UpdateHotel(ctx context.Context, hotel *service_models.Hotel) (*service_models.Hotel, error) {
	filter := bson.M{"_id": hotel.Id}
	update := bson.M{"$set": hotel}
	_, err := h.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, errors.New("failed to update hotel")
	}

	return hotel, nil
}

func NewHotelRepository(db *mongo.Database) HotelRepository {
	return &hotelRepository{
		collection: db.Collection("hotels"),
	}
}
