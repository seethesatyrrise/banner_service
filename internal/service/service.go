package service

import (
	"bannerService/internal/entity"
	"bannerService/internal/repo"
	"context"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

//type User interface {
//	CreateUser(ctx context.Context, user entity.User) (int, error)
//	UserById(ctx context.Context, id int) (entity.SegmentList, error)
//	AddDeleteSegment(ctx context.Context, segments entity.AddDelSegments) error
//	Operations(ctx context.Context, userOperations entity.UserOperations) ([]entity.Operation, error)
//}
//
//type Segment interface {
//	CreateSegment(ctx context.Context, segment entity.Segment) (int, error)
//	DeleteSegment(ctx context.Context, name string) error
//}

type Banner interface {
	CreateBanner(ctx context.Context, banner entity.Banner) (int, error)
	//FilterBanners() ()
	//UpdateBanner() ()
	//DeleteBanner() ()
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
