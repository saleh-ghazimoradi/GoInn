package cmd

import (
	"context"
	"fmt"
	"github.com/saleh-ghazimoradi/GoInn/config"
	gateway "github.com/saleh-ghazimoradi/GoInn/internal"
	"github.com/saleh-ghazimoradi/GoInn/logger"
	"github.com/saleh-ghazimoradi/GoInn/utils"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "runs the http server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("http called")

		// -------------------------------
		// Logger
		// -------------------------------
		log := logger.NewLogger(
			logger.WithLevel(slog.LevelDebug),
			logger.WithSource(true),
			logger.WithOutput(os.Stdout),
		)
		defer log.Close()

		defer log.Close()

		// -------------------------------
		// Config
		// -------------------------------
		cfg, err := config.GetConfig()
		if err != nil {
			log.Error("error loading config", slog.Any("err", err))
			return
		}

		// -------------------------------
		// Mongo (Go Driver v2 style)
		// -------------------------------
		mongo := utils.NewMongoDB(
			utils.WithHost(cfg.Mongo.Host),
			utils.WithPort(cfg.Mongo.Port),
			utils.WithDBName(cfg.Mongo.Name),
			utils.WithUser(cfg.Mongo.User),
			utils.WithAuthSource(cfg.Mongo.AuthSource),
			utils.WithPass(cfg.Mongo.Password),
			utils.WithMaxPoolSize(cfg.Mongo.MaxPoolSize),
			utils.WithMinPoolSize(cfg.Mongo.MinPoolSize),
			utils.WithTimeout(cfg.Mongo.Timeout),
		)

		client, db, err := mongo.Connect()
		if err != nil {
			log.Error("error connecting to mongo", slog.Any("err", err))
			return
		}
		log.Info("connected to MongoDB", slog.String("db", db.Name()))

		defer func() {
			if err := client.Disconnect(context.Background()); err != nil {
				log.Error("error disconnecting from mongo", slog.Any("err", err))
			}
		}()

		// -------------------------------
		// HTTP Server
		// -------------------------------
		server := gateway.NewServer(
			gateway.WithHost(cfg.Server.Host),
			gateway.WithPort(cfg.Server.Port),
			gateway.WithHandler(nil), // TODO: inject your handler/router
			gateway.WithReadTimeout(cfg.Server.ReadTimeout),
			gateway.WithWriteTimeout(cfg.Server.WriteTimeout),
			gateway.WithIdleTimeout(cfg.Server.IdleTimeout),
			gateway.WithTimeout(cfg.Server.Timeout),
		)

		if err := server.Connect(); err != nil {
			log.Error("server failed", slog.Any("err", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
