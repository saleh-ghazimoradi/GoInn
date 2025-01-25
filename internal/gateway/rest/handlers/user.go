package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/GoInn/internal/service"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

func (u *UserHandler) GetUserHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user, err := u.userService.GetUserById(context.Background(), id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func (u *UserHandler) GetUsersHandler(ctx *fiber.Ctx) error {
	users, err := u.userService.GetUsers(context.Background())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "success",
		"data":    users,
	})
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
