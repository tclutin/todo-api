package handler

import (
	"context"
	"log/slog"
	"net/http"
	note2 "todo/internal/service/note"
)

type NoteService interface {
	CreateNote(context.Context, note2.CreateNoteDTO) (note2.Note, error)
	UpdateNote(context.Context, note2.UpdateNoteDTO) (note2.Note, error)
	GetNoteByID(context.Context, uint64) (note2.Note, error)
	DeleteNote(ctx context.Context, uint642 uint64) error
	GetAllNotes(ctx context.Context) ([]note2.Note, error)
}

type Handler struct {
	logger      *slog.Logger
	noteService NoteService
}

func (h *Handler) InitNotesRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/note/create", h.CreateNote)
	router.HandleFunc("/note/get/", h.GetNoteByID)
	router.HandleFunc("/notes", h.GetAllNotes)
	router.HandleFunc("/note/delete/", h.DeleteNote)
	router.HandleFunc("/note/update/", h.UpdateNote)
	return router
}

func New(logger *slog.Logger, service NoteService) Handler {
	return Handler{
		logger:      logger,
		noteService: service,
	}
}
