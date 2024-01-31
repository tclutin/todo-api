package main

import (
	"log/slog"
	"os"
	"todo/internal/app"
	"todo/internal/config"
	"todo/internal/handler"
	"todo/internal/note"
	"todo/internal/repository"
)

func main() {
	//Initializing the config
	cfg := config.MustLoad()

	//Initializing the logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	//Initializing the noteRepo
	noteRepository := repository.NewNoteRepository(nil)

	//Initializing the note service
	noteService := note.NewService(logger, noteRepository)

	//Initializing the router
	router := handler.New(logger, noteService)

	//Initializing the app
	app := app.New(cfg, logger, router.InitNotesRoutes())

	app.Run()
}
