package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	Host string `env:"HOST" env-default:"localhost"`
	Port int    `env:"PORT" env-default:"25565"`
}

func MustLoad() *Config {
	var config Config
	var once sync.Once

	once.Do(func() {
		log.Println("Initial config of application")
		if err := cleanenv.ReadEnv(&config); err != nil {
			log.Fatalln("Error initial config of application:", err)
		}
	})
	return &config
}
