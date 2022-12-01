package main

import (
	"database/sql"
	"ticket/config"
	"ticket/repository"
	"ticket/server"
	"ticket/service"
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

	repo := repository.NewSqlRepository(db)

	svc := service.NewService(repo)

	svr := server.NewGinServer(svc, &cfg.Server)

	svr.Run()
}
