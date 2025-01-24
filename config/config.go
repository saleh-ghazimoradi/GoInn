package config

var AppConfig *Config

type Config struct {
	DBConfig DBConfig
	Server   ServerConfig
}

type DBConfig struct {
	DbHost     string `env:"DB_HOST,required"`
	DbPort     string `env:"DB_PORT,required"`
	DbUser     string `env:"DB_USER,required"`
	DbPassword string `env:"DB_PASSWORD,required"`
	DbName     string `env:"DB_NAME,required"`
}

type ServerConfig struct {
}

func LoadConfig() error {
	return nil
}
