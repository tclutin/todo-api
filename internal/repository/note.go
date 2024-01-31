package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"todo/internal/note"
)

type noteRepository struct {
	db *pgxpool.Pool
}

func (n noteRepository) Create(entity note.Note) error {

	return nil
}

func NewNoteRepository(database *pgxpool.Pool) *noteRepository {
	return &noteRepository{db: database}
}
