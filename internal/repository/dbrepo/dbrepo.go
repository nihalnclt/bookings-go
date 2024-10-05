package dbrepo

import (
	"database/sql"

	"github.com/nihalnclt/bookings-go/internal/config"
	"github.com/nihalnclt/bookings-go/internal/repository"
)

type postgresRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresRepo{
		App: a,
		DB:  conn,
	}
}
