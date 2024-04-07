package repo

import (
	"bannerService/internal/entity"
	"context"
	"github.com/jmoiron/sqlx"
)

type Banner interface {
	CreateBanner(ctx context.Context, banner entity.Banner) (int, error)
	//FilterBanners() ()
	//UpdateBanner() ()
	//DeleteBanner() ()
}

type UserBanner interface {
	//GetBanner(ctx context.Context, ) (int, error)
}

type Repository struct {
	Banner
	UserBanner
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Banner:     NewBanner(db),
		UserBanner: NewUserBanner(db),
	}
}
