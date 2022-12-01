package repository

import (
	"database/sql"
	"flight/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

//type IFlightRepo interface {
//	ListFlightsWithOffsetLimit(offset, limit int) []model.Flight
//	//TODO: offset and limit require full table scan, instead use: select * from tb where id>a limit b
//}

type Repo Querier

func NewSqlRepository(db *sql.DB) Repo {
	return New(db)
}

func NewSqlDatabase(cfg *config.DbConfig) (*sql.DB, error) {
	pool, err := sql.Open("pgx", cfg.Url)
	return pool, err
}
