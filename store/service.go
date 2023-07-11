package store

import (
	"context"
	"fmt"
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

func InitStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Init redis error: %v", err)
		return nil
	}

	fmt.Printf("\n Redis start success: pong msg = [%s]", pong)
	storeService.redisClient = redisClient

	return storeService
}

func SaveUrlMapping(short, original string) {
	err := storeService.redisClient.Set(ctx, short, original, cacheDuration).Err()
	if err != nil {
		fmt.Printf("Save key url error: %v", err)
	}
}

func GetOriginalUrl(short string) string {
	result, err := storeService.redisClient.Get(ctx, short).Result()
	if err != nil {
		fmt.Printf("Save key url error: %v", err)
		return ""
	}

	return result
}
