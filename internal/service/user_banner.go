package service

import (
	"bannerService/internal/entity"
	"bannerService/internal/repo"
	"context"
	"encoding/json"
)

type UserBannerService struct {
	repo repo.UserBanner
}

func NewUserBannerService(repo repo.UserBanner) *UserBannerService {
	return &UserBannerService{repo: repo}
}

func (s *UserBannerService) GetBanner(ctx context.Context, banner entity.UserBanner) (map[string]interface{}, error) {
	// todo add cache and check use_last_revision
	content, err := s.repo.GetBanner(ctx, banner.TagId, banner.FeatureId)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(content, &result)

	return result, err
}
