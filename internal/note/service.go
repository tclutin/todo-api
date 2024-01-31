package note

import (
	"context"
	"fmt"
	"log/slog"
)

type NoteRepo interface {
	Create() error
}

type service struct {
	logger *slog.Logger
	repo   NoteRepo
}

func (s service) Create(ctx context.Context, dto CreateNoteDTO) error {
	fmt.Println(dto)
	return nil
}

func (s service) Update(ctx context.Context, dto UpdateNoteDTO) error {

	return nil
}

func (s service) Delete(ctx context.Context, dto UpdateNoteDTO) error {

	return nil
}

func (s service) GetAll(ctx context.Context, dto UpdateNoteDTO) error {

	return nil
}

func (s service) GetByID(ctx context.Context, dto UpdateNoteDTO) error {

	return nil
}

func NewService(logger *slog.Logger) *service {
	return &service{logger: logger}
}
