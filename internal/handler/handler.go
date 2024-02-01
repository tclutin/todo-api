package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"todo/internal/note"
	custom "todo/pkg/http"
)

type NoteService interface {
	CreateNote(context.Context, note.CreateNoteDTO) (note.Note, error)
	UpdateNote(context.Context, note.UpdateNoteDTO) (note.Note, error)
	GetNoteByID(context.Context, uint64) (note.Note, error)
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

func (h *Handler) CreateNote(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var dto note.CreateNoteDTO
	if err := json.NewDecoder(request.Body).Decode(&dto); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	defer request.Body.Close()

	created, err := h.noteService.CreateNote(request.Context(), dto)
	if err != nil {
		h.logger.Error(err.Error())
		custom.SendJSON[string](writer, http.StatusBadRequest, err.Error())
		return
	}

	custom.SendJSON[note.Note](writer, http.StatusCreated, created)
	return
}

func (h *Handler) UpdateNote(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	if request.Method != http.MethodPatch {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var dto note.UpdateNoteDTO
	if err := json.NewDecoder(request.Body).Decode(&dto); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	defer request.Body.Close()

	updated, err := h.noteService.UpdateNote(request.Context(), dto)
	if err != nil {
		h.logger.Error(err.Error())
		custom.SendJSON[string](writer, http.StatusBadRequest, err.Error())
		return
	}

	custom.SendJSON[note.Note](writer, http.StatusOK, updated)
	return
}

func (h *Handler) GetNoteByID(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := request.URL.Query().Get("id")
	value, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(value)
}

func (h *Handler) GetAllNotes(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) DeleteNote(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
}

func New(logger *slog.Logger, service NoteService) Handler {
	return Handler{
		logger:      logger,
		noteService: service,
	}
}
