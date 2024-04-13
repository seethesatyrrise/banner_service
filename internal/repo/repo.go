package repo

import (
	"bannerService/internal/entity"
	"context"
	"github.com/jmoiron/sqlx"
)

type Banner interface {
	CreateBanner(ctx context.Context, banner entity.Banner) (int, error)
	FilterBanners(ctx context.Context, params entity.BannerFilters) ([]entity.BannerInfo, error)
	UpdateBanner(ctx context.Context, patch map[string]interface{}) error
}

type UserBanner interface {
	GetBanner(ctx context.Context, tagId, featureId int) ([]byte, error)
}

type BannerHistory interface {
	GetBannerHistory(ctx context.Context, id int) (entity.BannerHistory, error)
	SetBannerVersion(ctx context.Context, id, version int) error
}

type Deletion interface {
	DeleteFromDB(ctx context.Context, data entity.Deletion) error
}

type Repository struct {
	Banner
	UserBanner
	BannerHistory
	Deletion
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Banner:        NewBanner(db),
		UserBanner:    NewUserBanner(db),
		BannerHistory: NewBannerHistory(db),
		Deletion:      NewDeletion(db),
	}
}
