package main

import (
	"log/slog"
	"os"
	"todo/internal/config"
)

func main() {
	//Initializing the config
	cfg := config.MustLoad()

	//Initializing the logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

}
