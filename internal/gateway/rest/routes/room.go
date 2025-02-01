package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/GoInn/internal/gateway/rest/handlers"
)

func roomRoutes(app *fiber.App, handler *handlers.RoomHandler) {
	v1 := app.Group("/v1")
	v1.Post("/room", handler.CreateRoom)
	v1.Get("/room", handler.GetRooms)
	v1.Get("/hotel/:id/room", handler.GetRoomByHotelId)
}
