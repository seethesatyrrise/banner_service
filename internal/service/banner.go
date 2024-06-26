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

func (s *BannerService) FilterBanners(ctx context.Context, params entity.BannerFilters) ([]entity.BannerInfo, error) {
	return s.repo.FilterBanners(ctx, params)
}

func (s *BannerService) UpdateBanner(ctx context.Context, patch map[string]interface{}) error {
	return s.repo.UpdateBanner(ctx, patch)
}
