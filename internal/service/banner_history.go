package service

import (
	"bannerService/internal/entity"
	"bannerService/internal/repo"
	"context"
)

type BannerHistoryService struct {
	repo repo.BannerHistory
}

func NewBannerHistoryService(repo repo.BannerHistory) *BannerHistoryService {
	return &BannerHistoryService{repo: repo}
}

func (s *BannerHistoryService) GetBannerHistory(ctx context.Context, id int) (entity.BannerHistory, error) {
	return s.repo.GetBannerHistory(ctx, id)
}

func (s *BannerHistoryService) SetBannerVersion(ctx context.Context, id, version int) error {
	return s.repo.SetBannerVersion(ctx, id, version)
}
