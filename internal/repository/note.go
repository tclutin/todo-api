package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"log/slog"
	"strings"
	"todo/internal/note"
	"todo/pkg/client/postgresql"
)

type noteRepository struct {
	logger *slog.Logger
	client postgresql.Client
}

func (n *noteRepository) InitiDB() {
	sql := `
		CREATE TABLE IF NOT EXISTS todo (
    		id SERIAL PRIMARY KEY,
    		name TEXT NOT NULL,
    		content TEXT NOT NULL,
    		is_done BOOLEAN NOT NULL,
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)
    		`

	_, err := n.client.Exec(context.Background(), sql)
	if err != nil {
		log.Fatalln(err)
	}
}

func (n *noteRepository) CreateNote(ctx context.Context, entity note.Note) (note.Note, error) {
	sql := `
			INSERT INTO todo (name, content, is_done, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING *`

	row := n.client.QueryRow(ctx, sql, entity.Name, entity.Content, entity.IsDone, entity.CreatedAt, entity.UpdateAt)

	n.logger.Info("sql query:", strings.ReplaceAll(sql, "\t", ""))

	var note note.Note
	err := row.Scan(&note.ID, &note.Name, &note.Content, &note.IsDone, &note.CreatedAt, &note.UpdateAt)
	if err != nil {
		n.logger.Error("the query was unsuccessful", err)
		return note, err
	}
	n.logger.Info("the query was successful")
	return note, nil
}

func NewNoteRepository(database *pgxpool.Pool, logger *slog.Logger) *noteRepository {
	return &noteRepository{client: database, logger: logger}
}
