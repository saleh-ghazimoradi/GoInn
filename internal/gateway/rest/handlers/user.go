package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/GoInn/helper"
	"github.com/saleh-ghazimoradi/GoInn/internal/dto"
	"github.com/saleh-ghazimoradi/GoInn/internal/service"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

func (u *UserHandler) CreateUserHandler(ctx *fiber.Ctx) error {
	var user dto.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := helper.Validate(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	us, err := u.userService.CreateUser(ctx.Context(), &user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created",
		"user":    us,
	})
}

func (u *UserHandler) GetUserHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user, err := u.userService.GetUserById(ctx.Context(), id)
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
	users, err := u.userService.GetUsers(ctx.Context())
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
