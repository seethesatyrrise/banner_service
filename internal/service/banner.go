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
	id, err := s.repo.CreateBanner(ctx, banner)
	return id, err
}
