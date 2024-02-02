package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo/internal/note"
	custom "todo/pkg/http"
)

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
	writer.Header().Set("Content-Type", "application/json")
	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	value := request.URL.Query().Get("id")
	id, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	got, err := h.noteService.GetNoteByID(request.Context(), id)
	if err != nil {
		h.logger.Error(err.Error())
		custom.SendJSON[string](writer, http.StatusBadRequest, err.Error())
		return
	}
	custom.SendJSON[note.Note](writer, http.StatusOK, got)
	return
}

func (h *Handler) GetAllNotes(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	notes, err := h.noteService.GetAllNotes(request.Context())
	if err != nil {
		h.logger.Error(err.Error())
		custom.SendJSON[string](writer, http.StatusBadRequest, err.Error())
		return
	}

	custom.SendJSON[[]note.Note](writer, http.StatusOK, notes)
	return
}

func (h *Handler) DeleteNote(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	if request.Method != http.MethodDelete {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	value := request.URL.Query().Get("id")
	fmt.Println(value)
	id, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = h.noteService.DeleteNote(request.Context(), id); err != nil {
		h.logger.Error(err.Error())
		custom.SendJSON[string](writer, http.StatusBadRequest, err.Error())
		return
	}
	custom.SendJSON[string](writer, http.StatusOK, "Succesfully")
}
