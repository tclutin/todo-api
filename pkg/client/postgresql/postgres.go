package postgres

import "github.com/jackc/pgx"

type Client struct {
	db *pgx.ConnPool
}
