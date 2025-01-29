package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/GoInn/internal/dto"
	"github.com/saleh-ghazimoradi/GoInn/internal/service"
)

type HotelHandler struct {
	hotelService service.HotelService
}

func (h *HotelHandler) CreateHotel(ctx *fiber.Ctx) error {
	var req dto.CreateHotel
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	hotel, err := h.hotelService.CreateHotel(ctx.Context(), &req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "hotel successfully created",
		"data":    hotel,
	})
}

func NewHotelHandler(hotelService service.HotelService) *HotelHandler {
	return &HotelHandler{
		hotelService: hotelService,
	}
}
