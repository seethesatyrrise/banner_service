package service

import (
	"bannerService/internal/cache"
	"bannerService/internal/entity"
	"bannerService/internal/repo"
	"context"
	"encoding/json"
	"fmt"
)

type UserBannerService struct {
	repo  repo.UserBanner
	cache *cache.Cache
}

func NewUserBannerService(repo repo.UserBanner, cache *cache.Cache) *UserBannerService {
	return &UserBannerService{repo: repo, cache: cache}
}

func (s *UserBannerService) GetBanner(ctx context.Context, banner entity.UserBanner) (map[string]interface{}, error) {
	var content []byte
	var err error

	key := makeCacheKey(banner.FeatureId, banner.TagId)
	if !banner.UseLastRevision {
		content, err = s.cache.Cache.Get(ctx, key).Bytes()
	}
	if err != nil || banner.UseLastRevision {
		content, err = s.repo.GetBanner(ctx, banner.TagId, banner.FeatureId)
		if err != nil {
			return nil, err
		}
		s.cache.Cache.Del(ctx, key)
		err = s.cache.Cache.SetNX(ctx, key, content, cache.TTL).Err()
		if err != nil {
			return nil, err
		}
	}

	var result map[string]interface{}
	err = json.Unmarshal(content, &result)

	return result, err
}

func makeCacheKey(featureId, tagId int) string {
	return fmt.Sprintf("%d.%d", featureId, tagId)
}
