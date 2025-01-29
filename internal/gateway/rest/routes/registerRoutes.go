package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/GoInn/internal/gateway/rest/handlers"
	"github.com/saleh-ghazimoradi/GoInn/internal/repository"
	"github.com/saleh-ghazimoradi/GoInn/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(app *fiber.App, db *mongo.Database) {
	health := handlers.NewHealthHandler()
	userRepository := repository.NewUserRepository(db)
	hotelRepository := repository.NewHotelRepository(db)
	hotelService := service.NewHotelService(hotelRepository)
	hotelHandler := handlers.NewHotelHandler(hotelService)
	userService := service.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	healthRoutes(app, health)
	userRoutes(app, userHandler)
	hotelRoutes(app, hotelHandler)
}
