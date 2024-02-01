package note

import (
	"context"
	"errors"
	"log/slog"
	"time"
)

type NoteRepository interface {
	CreateNote(context.Context, Note) (Note, error)
	UpdateNote(context.Context, Note) (Note, error)
	GetNoteById(context.Context, uint64) (Note, error)
}

type service struct {
	logger *slog.Logger
	repo   NoteRepository
}

func (s *service) CreateNote(ctx context.Context, dto CreateNoteDTO) (Note, error) {
	if !dto.Validate() {
		return Note{}, errors.New("fields must have values")
	}

	if len(dto.Name) > 30 || len(dto.Name) < 3 {
		return Note{}, errors.New("field name must be 3<x<10")
	}

	if len(dto.Content) > 20 {
		return Note{}, errors.New("field content must be x<20")
	}

	model := Note{
		Name:      dto.Name,
		Content:   dto.Content,
		IsDone:    false,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	note, err := s.repo.CreateNote(ctx, model)
	if err != nil {
		return note, errors.New("internal database error")
	}

	return note, nil
}

func (s *service) UpdateNote(ctx context.Context, dto UpdateNoteDTO) (Note, error) {
	if len(dto.Name) > 30 || len(dto.Name) < 3 {
		return Note{}, errors.New("field name must be 3<x<10")
	}

	if len(dto.Content) > 20 {
		return Note{}, errors.New("field content must be x<20")
	}

	noteById, err := s.repo.GetNoteById(ctx, dto.ID)
	if err != nil {
		return noteById, errors.New("internal database error")
	}

	if dto.Content != "" {
		noteById.Content = dto.Content
	}

	if dto.Name != "" {
		noteById.Name = dto.Name
	}

	if dto.IsDone != noteById.IsDone {
		noteById.IsDone = dto.IsDone
	}

	noteById.UpdateAt = time.Now()

	updatedNote, err := s.repo.UpdateNote(ctx, noteById)
	if err != nil {
		return noteById, errors.New("internal database error")
	}

	return updatedNote, nil
}

func (s *service) Delete(ctx context.Context, dto UpdateNoteDTO) error {

	return nil
}

func (s *service) GetAll(ctx context.Context, dto UpdateNoteDTO) error {

	return nil
}

func (s *service) GetNoteByID(ctx context.Context, id uint64) (Note, error) {

	return Note{}, nil
}

func NewService(logger *slog.Logger, repo NoteRepository) *service {
	return &service{logger: logger, repo: repo}
}
