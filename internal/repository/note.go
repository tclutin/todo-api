package repository

import (
	"database/sql"
	"todo/internal/note"
)

type noteRepository struct {
	db *sql.DB
}

func (n noteRepository) Create(entity note.Note) error {

	return nil
}

func NewNoteRepository(database *sql.DB) *noteRepository {
	return &noteRepository{db: database}
}
