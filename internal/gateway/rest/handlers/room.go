package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/GoInn/internal/dto"
	"github.com/saleh-ghazimoradi/GoInn/internal/service"
	"net/http"
)

type RoomHandler struct {
	roomService service.RoomService
}

func (r *RoomHandler) CreateRoom(ctx *fiber.Ctx) error {
	var room dto.CreateRoom
	if err := ctx.BodyParser(&room); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ro, err := r.roomService.CreateRoom(ctx.Context(), &room)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "room created",
		"room":    ro,
	})
}

func (r *RoomHandler) GetRooms(ctx *fiber.Ctx) error {
	rooms, err := r.roomService.GetRooms(ctx.Context())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "rooms",
		"rooms":   rooms,
	})
}

func (r *RoomHandler) GetRoomByHotelId(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	rooms, err := r.roomService.GetRoomByHotelId(ctx.Context(), id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "hotel rooms",
		"rooms":   rooms,
	})
}

func NewRoomHandler(roomService service.RoomService) *RoomHandler {
	return &RoomHandler{
		roomService: roomService,
	}
}
