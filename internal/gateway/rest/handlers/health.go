package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/GoInn/config"
)

type HealthHandler struct {
}

func (h *HealthHandler) Health(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"env":     config.AppConfig.ServerConfig.Port,
		"version": config.AppConfig.ServerConfig.Version,
	})
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}
