package service

import (
	"bannerService/internal/entity"
	"bannerService/internal/repo"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type BannerService struct {
	repo repo.Banner
}

func NewBannerService(repo repo.Banner) *BannerService {
	return &BannerService{repo: repo}
}

func (s *BannerService) CreateBanner(ctx context.Context, banner entity.Banner) (int, error) {
	content, err := json.Marshal(banner.Content)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("BannerService.CreateBanner: %s", err.Error()))
	}

	bannerId, err := s.repo.AddBanner(ctx, entity.BannerToDB{
		Content:  content,
		IsActive: banner.IsActive,
	})
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("BannerService.CreateBanner: %s", err.Error()))
	}

	bannerRelations := make([]entity.BannerRelationsToDB, 0, len(banner.TagIds))
	for _, tagId := range banner.TagIds {
		bannerRelations = append(bannerRelations, entity.BannerRelationsToDB{
			TagId:     tagId,
			FeatureId: banner.FeatureId,
			BannerId:  bannerId,
		})
	}

	if err = s.repo.AddBannerRelations(ctx, bannerRelations); err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("BannerService.CreateBanner: %s", err.Error()))
	}

	return bannerId, err
}

func (s *BannerService) DeleteBanner(ctx context.Context, id int) error {
	return s.repo.DeleteBanner(ctx, id)
}
