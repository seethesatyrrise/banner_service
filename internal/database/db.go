package database

import (
	"bannerService/internal/config"
	"bannerService/internal/utils"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	DB *sqlx.DB
}

func New(cfg *config.DB) (*DB, error) {

	db, err := sqlx.Open("postgres", dbUrl(*cfg))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func dbUrl(cfg config.DB) string {
	utils.Logger.Info(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name))
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
}
