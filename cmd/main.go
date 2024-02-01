package main

import (
	"context"
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
	pgxPool := postgresql.NewClient(context.Background(), cfg.DB)

	//Initializing the noteRepo
	noteRepository := repository.NewNoteRepository(pgxPool, logger)
	noteRepository.InitiDB()

	//Initializing the note service
	noteService := note.NewService(logger, noteRepository)

	//Initializing the router
	router := handler.New(logger, noteService)

	//Initializing the app
	app := app.New(cfg, logger, router.InitNotesRoutes())

	//start app
	app.Run()
}
