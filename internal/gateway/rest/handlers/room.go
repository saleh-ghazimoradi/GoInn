package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/GoInn/internal/service"
)

type RoomHandler struct {
	roomService service.RoomService
}

func (r *RoomHandler) CreateRoom(ctx *fiber.Ctx) error {
	return nil
}

func NewRoomHandler(roomService service.RoomService) *RoomHandler {
	return &RoomHandler{
		roomService: roomService,
	}
}
