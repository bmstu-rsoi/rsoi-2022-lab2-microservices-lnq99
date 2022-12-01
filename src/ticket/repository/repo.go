package repository

import (
	"database/sql"
	"ticket/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Repo Querier

func NewSqlRepository(db *sql.DB) Repo {
	return New(db)
}

func NewSqlDatabase(cfg *config.DbConfig) (*sql.DB, error) {
	pool, err := sql.Open("pgx", cfg.Url)
	return pool, err
}
