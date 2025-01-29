package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/GoInn/internal/gateway/rest/handlers"
)

func hotelRoutes(app *fiber.App, hotelHandler *handlers.HotelHandler) {
	v1 := app.Group("/v1")
	v1.Post("/hotel", hotelHandler.CreateHotel)
}
