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
	DB   string `env:"DB" env-required:"true"`
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
