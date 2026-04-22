package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"go.rest.api/internal/config"
)

var db *sql.DB

func Init(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	db = conn
	return db, nil
}
