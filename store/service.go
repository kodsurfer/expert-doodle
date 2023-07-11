package store

import (
	"context"
	"fmt"
	redis "github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

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
