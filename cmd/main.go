package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"todo/internal/app"
	"todo/internal/config"
	"todo/internal/handler"
	"todo/internal/note"
	"todo/internal/repository"
	"todo/pkg/client/postgresql"
)

func main() {
	//Initializing the config
	cfg := config.MustLoad()

	//Initializing the logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	//Initializing pgxpool
	connToPGQ := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.DBConfig.Username, cfg.DBConfig.Password, cfg.DBConfig.Host, cfg.DBConfig.Port, cfg.DBConfig.Database)
	pgxPool := postgresql.NewClient(context.Background(), connToPGQ)

	//Initializing the noteRepo
	noteRepository := repository.NewNoteRepository(pgxPool)

	//Initializing the note service
	noteService := note.NewService(logger, noteRepository)

	//Initializing the router
	router := handler.New(logger, noteService)

	//Initializing the app
	app := app.New(cfg, logger, router.InitNotesRoutes())

	//start app
	app.Run()
}
