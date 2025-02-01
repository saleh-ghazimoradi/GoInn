package service

import (
	"context"
	"fmt"
	"github.com/saleh-ghazimoradi/GoInn/internal/dto"
	"github.com/saleh-ghazimoradi/GoInn/internal/repository"
	"github.com/saleh-ghazimoradi/GoInn/internal/service/service_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomService interface {
	CreateRoom(ctx context.Context, input *dto.CreateRoom) (*service_models.Room, error)
	GetRooms(ctx context.Context) ([]*service_models.Room, error)
	GetRoomByHotelId(ctx context.Context, id string) ([]*service_models.Room, error)
}

type roomService struct {
	roomRepository  repository.RoomRepository
	hotelRepository repository.HotelRepository
}

func (r *roomService) CreateRoom(ctx context.Context, input *dto.CreateRoom) (*service_models.Room, error) {

	room, err := r.roomRepository.InsertRoom(ctx, &service_models.Room{
		Type:      service_models.RoomType(input.Type),
		BasePrice: input.BasePrice,
		Price:     input.Price,
		HotelId:   input.HotelId,
	})
	if err != nil {
		return nil, err
	}

	hotel, err := r.hotelRepository.GetHotelById(ctx, input.HotelId)
	if err != nil {
		return nil, err
	}

	hotel.Rooms = append(hotel.Rooms, room.Id)

	_, err = r.hotelRepository.UpdateHotel(ctx, hotel)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (r *roomService) GetRooms(ctx context.Context) ([]*service_models.Room, error) {
	return r.roomRepository.GetRooms(ctx)
}

func (r *roomService) GetRoomByHotelId(ctx context.Context, id string) ([]*service_models.Room, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid hotel ID: %v", err)
	}

	rooms, err := r.roomRepository.GetRoomsByHotelId(ctx, oid)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch rooms: %v", err)
	}

	if len(rooms) == 0 {
		return nil, fmt.Errorf("no rooms found for hotel ID: %s", id)
	}

	return rooms, nil
}

func NewRoomService(roomRepository repository.RoomRepository, hotelRepository repository.HotelRepository) RoomService {
	return &roomService{
		roomRepository:  roomRepository,
		hotelRepository: hotelRepository,
	}
}
