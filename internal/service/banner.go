package service

import (
	"bannerService/internal/entity"
	"bannerService/internal/repo"
	"context"
)

type BannerService struct {
	repo repo.Banner
}

func NewBannerService(repo repo.Banner) *BannerService {
	return &BannerService{repo: repo}
}

func (s *BannerService) CreateBanner(ctx context.Context, banner entity.Banner) (int, error) {
	return s.repo.CreateBanner(ctx, banner)
}

func (s *BannerService) DeleteBanner(ctx context.Context, id int) error {
	return s.repo.DeleteBanner(ctx, id)
}

func (s *BannerService) FilterBanners(ctx context.Context, params entity.BannerFilters) ([]entity.BannerInfo, error) {
	return s.repo.FilterBanners(ctx, params)
}
