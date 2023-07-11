package store

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const cacheDuration = 1 * time.Hour
