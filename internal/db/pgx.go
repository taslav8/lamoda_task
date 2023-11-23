package db

import (
	"fmt"
	"log"

	"lamoda_task/internal/config"

	"github.com/jmoiron/sqlx"
)

func NewPgx(cfg *config.Db) (*sqlx.DB, error) {
	log.Printf("initialized master and slave DSNs: %v", cfg.GetDSN())

	db, err := sqlx.Connect("postgres", cfg.GetDSN())
	if err != nil {
		return nil, fmt.Errorf("got error while connecting to database: %w", err)
	}
	log.Printf("connected to database")

	db.SetMaxOpenConns(cfg.DBMaxOpenConns)
	db.SetMaxIdleConns(cfg.DBMaxIdleConns)
	db.SetConnMaxLifetime(cfg.DBConnMaxLifetime)
	log.Printf("set database config")

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("got error while pinging database: %w", err)
	}
	log.Printf("successfully connected to database")
	return db, nil
}
