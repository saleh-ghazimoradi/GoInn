package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/GoInn/internal/gateway/rest/handlers"
)

func userRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	v1 := app.Group("/v1")
	v1.Get("/user", userHandler.GetUsersHandler)
	v1.Get("/user/:id", userHandler.GetUserHandler)
}
