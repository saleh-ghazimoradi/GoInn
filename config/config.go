package config

import (
	"github.com/caarlos0/env/v11"
	"sync"
	"time"
)

var (
	appConfig *Config
	once      sync.Once
	initErr   error
)

type Config struct {
	Server      Server
	Mongo       Mongo
	Application Application
}

type Server struct {
	Host         string        `env:"SERVER_HOST"`
	Port         string        `env:"SERVER_PORT"`
	ReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT"`
	WriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT"`
	IdleTimeout  time.Duration `env:"SERVER_IDLE_TIMEOUT"`
	Timeout      time.Duration `env:"SERVER_TIMEOUT"`
}

type Mongo struct {
	Host        string        `env:"MONGODB_HOST"`
	Port        int           `env:"MONGODB_PORT"`
	User        string        `env:"MONGODB_USER"`
	Password    string        `env:"MONGODB_PASSWORD"`
	Name        string        `env:"MONGODB_NAME"`
	AuthSource  string        `env:"MONGODB_AUTH_SOURCE"`
	MaxPoolSize uint64        `env:"MONGODB_MAX_POOL_SIZE"`
	MinPoolSize uint64        `env:"MONGODB_MIN_POOL_SIZE"`
	Timeout     time.Duration `env:"MONGODB_TIMEOUT"`
}

type Application struct {
	Env     string `env:"APPLICATION_ENV"`
	Version string `env:"APPLICATION_VERSION"`
}

func GetConfig() (*Config, error) {
	once.Do(func() {
		appConfig = &Config{}
		initErr = env.Parse(appConfig)
		if initErr != nil {
			appConfig = nil
		}
	})
	return appConfig, initErr
}
