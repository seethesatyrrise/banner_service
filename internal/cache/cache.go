package cache

import (
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Cache *redis.Client
}

func New() *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Cache{rdb}
}
