package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/saleh-ghazimoradi/GoInn/slg"
	"time"
)

var AppConfig *Config

type Config struct {
	ServerConfig ServerConfig
	MongoConfig  MongoConfig
}

type ServerConfig struct {
	BodyLimit    int           `env:"BODY_LIMIT"`    // 1024 * 1024
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT"` // 10s
	ReadTimeout  time.Duration `env:"READ_TIMEOUT"`  // 5s
	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT"`  // 30s
	RateLimit    int           `env:"RATE_LIMIT"`    // 100
	RateLimitExp time.Duration `env:"RATE_EXP"`      // 60s
	Port         string        `env:"PORT"`          // 3000
	Timeout      time.Duration `env:"TIMEOUT"`       // 30s
}

type MongoConfig struct {
	DbHost     string        `env:"DB_HOST,required"`
	DbPort     string        `env:"DB_PORT,required"`
	DbUser     string        `env:"DB_USER,required"`
	DbPassword string        `env:"DB_PASSWORD,required"`
	DbName     string        `env:"DB_NAME,required"`
	DbTimeout  time.Duration `env:"DB_TIMEOUT,required"`
	DbUri      string        `env:"DB_URI,required"`
}

func LoadConfig() error {
	config := &Config{}

	if err := env.Parse(config); err != nil {
		slg.Logger.Error("error loading config", "error", err)
		return err
	}

	AppConfig = config

	return nil
}
