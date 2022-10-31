package main

import (
	"app/flight/config"
	"app/flight/repository"
	"app/flight/server"
	"app/flight/service"
	"database/sql"
)

func main() {
	var err error
	var db *sql.DB
	var cfg *config.Config

	if cfg, err = config.LoadConfig(); err != nil {
		panic(err)
	}

	if db, err = repository.NewSqlDatabase(&cfg.Db); err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	repo := repository.NewSqlRepository(db)

	svc := service.NewService(repo)

	svr := server.NewGinServer(svc, &cfg.Server)

	svr.Run()
}
