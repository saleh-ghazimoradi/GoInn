/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/saleh-ghazimoradi/GoInn/config"
	"github.com/saleh-ghazimoradi/GoInn/internal/repository"
	"github.com/saleh-ghazimoradi/GoInn/internal/service"
	"github.com/saleh-ghazimoradi/GoInn/logger"
	"github.com/saleh-ghazimoradi/GoInn/utils"

	"github.com/spf13/cobra"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seeding mongodb",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("seed called")
		client, database, err := utils.ConnectToMongoDB(config.AppConfig.DbConfig.DbUri, config.AppConfig.DbConfig.DbName)
		if err != nil {
			logger.Logger.Error("failed to connect to mongodb")
		}
		defer func() {
			if err := client.Disconnect(context.Background()); err != nil {
				logger.Logger.Error("error disconnecting from mongodb")
			}
		}()

		hotelRepository := repository.NewHotelRepository(database)
		roomRepository := repository.NewRoomRepository(database)
		userRepository := repository.NewUserRepository(database)
		seed := service.NewSeedService(roomRepository, hotelRepository, userRepository)
		if err := seed.Seed(context.Background()); err != nil {
			logger.Logger.Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
