package service

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/GoInn/internal/dto"
	"github.com/saleh-ghazimoradi/GoInn/internal/repository"
	"github.com/saleh-ghazimoradi/GoInn/internal/service/service_models"
)

type HotelService interface {
	CreateHotel(ctx context.Context, hotel *dto.CreateHotel) (*service_models.Hotel, error)
}

type hotelService struct {
	hotelRepository repository.HotelRepository
}

func (h *hotelService) CreateHotel(ctx context.Context, hotel *dto.CreateHotel) (*service_models.Hotel, error) {
	hot, err := h.hotelRepository.InsertHotel(ctx, &service_models.Hotel{
		Name:     hotel.Name,
		Location: hotel.Location,
		Rooms:    hotel.Rooms,
	})
	if err != nil {
		return nil, errors.New("error on creating hotel")
	}
	return hot, nil
}

func NewHotelService(hotelRepository repository.HotelRepository) HotelService {
	return &hotelService{
		hotelRepository: hotelRepository,
	}
}
