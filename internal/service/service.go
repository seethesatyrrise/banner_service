package service

import (
	"bannerService/internal/entity"
	"bannerService/internal/repo"
	"context"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Banner interface {
	CreateBanner(ctx context.Context, banner entity.Banner) (int, error)
	FilterBanners(ctx context.Context, params entity.BannerFilters) ([]entity.BannerInfo, error)
	UpdateBanner(ctx context.Context, patch map[string]interface{}) error
	DeleteBanner(ctx context.Context, id int) error
}

type UserBanner interface {
	GetBanner(ctx context.Context, banner entity.UserBanner) (map[string]interface{}, error)
}

type Service struct {
	Banner
	UserBanner
}

func New(repo *repo.Repository) *Service {
	return &Service{
		Banner:     NewBannerService(repo),
		UserBanner: NewUserBannerService(repo),
	}
}
