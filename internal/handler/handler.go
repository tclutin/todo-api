package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"todo/internal/note"
	custom "todo/pkg/http"
)

type NoteService interface {
	Create(context.Context, note.CreateNoteDTO) (note.Note, error)
	Update(context.Context, note.UpdateNoteDTO) error
}

type Handler struct {
	logger      *slog.Logger
	noteService NoteService
}

func (h Handler) InitNotesRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/note/create", h.CreateNote)
	router.HandleFunc("/note/get/", h.GetNote)
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

	ob, err := h.noteService.Create(request.Context(), dto)
	if err != nil {
		h.logger.Error(err.Error())
		custom.SendJSON[string](writer, http.StatusBadRequest, err.Error())
		return
	}

	custom.SendJSON[note.Note](writer, http.StatusCreated, ob)
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

	if err := h.noteService.Update(request.Context(), dto); err != nil {
		custom.SendJSON[string](writer, http.StatusBadRequest, err.Error())
		return
	}

	custom.SendJSON[string](writer, http.StatusOK, "Successfully")
	return
}

func (h *Handler) GetNote(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := request.URL.Query().Get("id")
	fmt.Printf("%T\n", id)
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
