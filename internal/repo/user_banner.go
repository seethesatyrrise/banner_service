package repo

import (
	"github.com/jmoiron/sqlx"
)

type UserBannerRepo struct {
	db *sqlx.DB
}

func NewUserBanner(db *sqlx.DB) *UserBannerRepo {
	return &UserBannerRepo{db: db}
}
