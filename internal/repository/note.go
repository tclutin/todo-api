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

func (n *noteRepository) DeleteNote(ctx context.Context, id uint64) error {
	sql := `DELETE FROM todo WHERE id = $1`
	_, err := n.client.Exec(ctx, sql, id)
	if err != nil {
		return err
	}
	return nil
}

func (n *noteRepository) GetNoteById(ctx context.Context, id uint64) (note.Note, error) {
	sql := `
			SELECT * FROM todo
			WHERE id = $1
			
			`

	n.logger.Info("sql query:", strings.ReplaceAll(sql, "\t", ""))

	row := n.client.QueryRow(ctx, sql, id)

	var note note.Note

	err := row.Scan(&note.ID, &note.Name, &note.Content, &note.IsDone, &note.CreatedAt, &note.UpdateAt)
	if err != nil {
		n.logger.Error("the query was unsuccessful", err)
		return note, err
	}
	n.logger.Info("the query was successful")
	return note, nil
}

func (n *noteRepository) UpdateNote(ctx context.Context, entity note.Note) (note.Note, error) {
	sql := `
			UPDATE todo
			SET name = $1, content = $2, is_done = $3, updated_at = $4
			WHERE id = $5
			RETURNING *`

	row := n.client.QueryRow(ctx, sql, entity.Name, entity.Content, entity.IsDone, entity.UpdateAt, entity.ID)

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

func (n *noteRepository) GetAllNotes(ctx context.Context) ([]note.Note, error) {
	sql := `SELECT * FROM todo`

	n.logger.Info("sql query:", strings.ReplaceAll(sql, "\t", ""))

	rows, err := n.client.Query(ctx, sql)
	if err != nil {
		n.logger.Error("the query was unsuccessful", err)
		return nil, err
	}

	notes := make([]note.Note, 0)

	for rows.Next() {
		var note note.Note

		err = rows.Scan(&note.ID, &note.Name, &note.Content, &note.IsDone, &note.CreatedAt, &note.UpdateAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	n.logger.Info("the query was successful")
	return notes, err
}

func NewNoteRepository(database *pgxpool.Pool, logger *slog.Logger) *noteRepository {
	return &noteRepository{client: database, logger: logger}
}
