package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/GoInn/internal/dto"
	"github.com/saleh-ghazimoradi/GoInn/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
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

func (h *HotelHandler) GetHotelById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	hotel, err := h.hotelService.GetHotelById(ctx.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "hotel does not exist",
			})
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    hotel,
	})
}

func (h *HotelHandler) GetHotels(ctx *fiber.Ctx) error {

	hotels, err := h.hotelService.GetHotels(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "success",
		"data":    hotels,
	})
}

func NewHotelHandler(hotelService service.HotelService) *HotelHandler {
	return &HotelHandler{
		hotelService: hotelService,
	}
}
