package handler

import (
	"context"
	"log/slog"
	"net/http"
	"todo/internal/note"
)

type NoteService interface {
	CreateNote(context.Context, note.CreateNoteDTO) (note.Note, error)
	UpdateNote(context.Context, note.UpdateNoteDTO) (note.Note, error)
	GetNoteByID(context.Context, uint64) (note.Note, error)
	DeleteNote(ctx context.Context, uint642 uint64) error
	GetAllNotes(ctx context.Context) ([]note.Note, error)
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
