package service

import "bannerService/internal/repo"

type UserBannerService struct {
	repo repo.UserBanner
}

func NewUserBannerService(repo repo.UserBanner) *UserBannerService {
	return &UserBannerService{repo: repo}
}
