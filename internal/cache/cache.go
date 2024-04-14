package cache

import (
	"bannerService/internal/config"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const TTL = 5 * time.Minute

type Cache struct {
	Cache *redis.Client
}

func New(cfg *config.Cache) *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Cache{rdb}
}
