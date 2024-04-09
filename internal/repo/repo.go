package repo

import (
	"bannerService/internal/entity"
	"context"
	"github.com/jmoiron/sqlx"
)

type Banner interface {
	CreateBanner(ctx context.Context, banner entity.Banner) (int, error)
	FilterBanners(ctx context.Context, params entity.BannerFilters) ([]entity.BannerInfo, error)
	//UpdateBanner() ()
	DeleteBanner(ctx context.Context, id int) error
}

type UserBanner interface {
	GetBanner(ctx context.Context, tagId, featureId int) ([]byte, error)
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
