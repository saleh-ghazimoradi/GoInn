package internal

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/GoInn/config"
	"github.com/saleh-ghazimoradi/GoInn/internal/gateway/rest/routes"
	"github.com/saleh-ghazimoradi/GoInn/logger"
	"github.com/saleh-ghazimoradi/GoInn/utils"
	"log"
)

func Server() error {
	app := fiber.New()

	client, database, err := utils.ConnectToMongoDB(config.AppConfig.DbConfig.DbUri, config.AppConfig.DbConfig.DbName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting MongoDB: %v", err)
		}
	}()

	routes.RegisterRoutes(app, database)

	logger.Logger.Info("starting server", "addr", config.AppConfig.ServerConfig.Port, "env", config.AppConfig.ServerConfig.ENV)

	if err := app.Listen(config.AppConfig.ServerConfig.Port); err != nil {
		return err
	}

	return nil
}
