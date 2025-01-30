package service

import (
	"context"
	"github.com/saleh-ghazimoradi/GoInn/internal/dto"
	"github.com/saleh-ghazimoradi/GoInn/internal/repository"
	"github.com/saleh-ghazimoradi/GoInn/internal/service/service_models"
)

type RoomService interface {
	CreateRoom(ctx context.Context, input *dto.CreateRoom) (*service_models.Room, error)
}

type roomService struct {
	roomRepository  repository.RoomRepository
	hotelRepository repository.HotelRepository
}

func (r *roomService) CreateRoom(ctx context.Context, input *dto.CreateRoom) (*service_models.Room, error) {
	return r.roomRepository.InsertRoom(ctx, &service_models.Room{
		Type:      service_models.SingleRoomType,
		BasePrice: input.BasePrice,
	})
}

func NewRoomService(roomRepository repository.RoomRepository) RoomService {
	return &roomService{
		roomRepository: roomRepository,
	}
}
