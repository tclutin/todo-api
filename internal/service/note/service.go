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
	DeleteNote(context.Context, uint64) error
	GetAllNotes(context.Context) ([]Note, error)
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

func (s *service) DeleteNote(ctx context.Context, id uint64) error {
	_, err := s.GetNoteByID(ctx, id)
	if err != nil {
		return errors.New("the note was not found")
	}

	if err = s.repo.DeleteNote(ctx, id); err != nil {
		return errors.New("internal database error")
	}
	return nil
}

func (s *service) GetAllNotes(ctx context.Context) ([]Note, error) {
	notes, err := s.repo.GetAllNotes(ctx)
	if err != nil {
		return nil, errors.New("internal database error")
	}

	return notes, err
}

func (s *service) GetNoteByID(ctx context.Context, id uint64) (Note, error) {
	note, err := s.repo.GetNoteById(ctx, id)
	if err != nil {
		return Note{}, errors.New("internal database error")
	}
	return note, nil
}

func NewService(logger *slog.Logger, repo NoteRepository) *service {
	return &service{logger: logger, repo: repo}
}
