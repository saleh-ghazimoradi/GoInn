package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/GoInn/internal/gateway/rest/handlers"
)

func healthRoutes(app *fiber.App, handler *handlers.HealthHandler) {
	app.Get("/v1/health", handler.Health)
}
