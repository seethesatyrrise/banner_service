package service

import (
	"bannerService/internal/cache"
	"bannerService/internal/deletion"
	"bannerService/internal/entity"
	"bannerService/internal/repo"
	"context"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Banner interface {
	CreateBanner(ctx context.Context, banner entity.Banner) (int, error)
	FilterBanners(ctx context.Context, params entity.BannerFilters) ([]entity.BannerInfo, error)
	UpdateBanner(ctx context.Context, patch map[string]interface{}) error
}

type UserBanner interface {
	GetBanner(ctx context.Context, banner entity.UserBanner) (map[string]interface{}, error)
}

type BannerHistory interface {
	GetBannerHistory(ctx context.Context, id int) (entity.BannerHistory, error)
	SetBannerVersion(ctx context.Context, id, version int) error
}

type Deletion interface {
	AddFeatureToDeletionQueue(ctx context.Context, featureId int)
	AddTagToDeletionQueue(ctx context.Context, tagId int)
	AddIdToDeletionQueue(ctx context.Context, id int)
	Close() error
}

type Service struct {
	Banner
	UserBanner
	BannerHistory
	Deletion
}

func New(repo *repo.Repository, cache *cache.Cache, queue *deletion.DeletionQueue) *Service {
	return &Service{
		Banner:        NewBannerService(repo),
		UserBanner:    NewUserBannerService(repo, cache),
		BannerHistory: NewBannerHistoryService(repo),
		Deletion:      NewDeletionService(repo, queue),
	}
}

func (s *Service) Close() error {
	return s.Deletion.Close()
}
