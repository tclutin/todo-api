package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

type Config struct {
	Host string `env:"HOST" env-default:"localhost"`
	Port string `env:"PORT" env-default:"25565"`
	DBConfig
}

type DBConfig struct {
	Host     string `env:"DB_HOST" env-required:"true"`
	Port     string `env:"DB_PORT" env-required:"true"`
	Username string `env:"DB_USERNAME" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
	Database string `env:"DB_NAME" env-required:"true"`
}

func MustLoad() *Config {
	var config Config
	var once sync.Once

	once.Do(func() {
		log.Println("Initial config of application")
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found")
		}

		if err := cleanenv.ReadEnv(&config); err != nil {
			log.Fatalln("Error initial config of application:", err)
		}
	})
	return &config
}
