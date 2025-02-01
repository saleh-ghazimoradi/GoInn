package service

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/GoInn/internal/dto"
	"github.com/saleh-ghazimoradi/GoInn/internal/repository"
	"github.com/saleh-ghazimoradi/GoInn/internal/service/service_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelService interface {
	CreateHotel(ctx context.Context, hotel *dto.CreateHotel) (*service_models.Hotel, error)
	GetHotels(ctx context.Context) ([]*service_models.Hotel, error)
	GetHotelById(ctx context.Context, id string) (*service_models.Hotel, error)
	UpdateHotel(ctx context.Context, hotel *service_models.Hotel) (*service_models.Hotel, error)
}

type hotelService struct {
	hotelRepository repository.HotelRepository
}

func (h *hotelService) CreateHotel(ctx context.Context, hotel *dto.CreateHotel) (*service_models.Hotel, error) {
	hot, err := h.hotelRepository.InsertHotel(ctx, &service_models.Hotel{
		Name:     hotel.Name,
		Location: hotel.Location,
		Rooms:    []primitive.ObjectID{},
		Rating:   hotel.Rating,
	})
	if err != nil {
		return nil, errors.New("error on creating hotel")
	}
	return hot, nil
}

func (h *hotelService) GetHotels(ctx context.Context) ([]*service_models.Hotel, error) {
	return h.hotelRepository.GetHotels(ctx)
}

func (h *hotelService) GetHotelById(ctx context.Context, id string) (*service_models.Hotel, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}
	return h.hotelRepository.GetHotelById(ctx, oid)
}

func (h *hotelService) UpdateHotel(ctx context.Context, hotel *service_models.Hotel) (*service_models.Hotel, error) {
	return nil, nil
}

func NewHotelService(hotelRepository repository.HotelRepository) HotelService {
	return &hotelService{
		hotelRepository: hotelRepository,
	}
}
