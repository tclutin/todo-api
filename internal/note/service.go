package note

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
)

type NoteRepository interface {
	CreateNote(context.Context, Note) (Note, error)
}

type service struct {
	logger *slog.Logger
	repo   NoteRepository
}

func (s *service) CreateNote(ctx context.Context, dto CreateNoteDTO) (Note, error) {
	if !dto.Validate() {
		return Note{}, errors.New("fields must have values")
	}

	if len(dto.Name) > 10 || len(dto.Name) < 3 {
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

func (s *service) Update(ctx context.Context, dto UpdateNoteDTO) error {
	fmt.Printf("%+v", dto)
	return errors.New("пошёл нахуй")
}

func (s *service) Delete(ctx context.Context, dto UpdateNoteDTO) error {

	return nil
}

func (s *service) GetAll(ctx context.Context, dto UpdateNoteDTO) error {

	return nil
}

func (s *service) GetByID(ctx context.Context, dto UpdateNoteDTO) error {

	return nil
}

func NewService(logger *slog.Logger, repo NoteRepository) *service {
	return &service{logger: logger, repo: repo}
}
