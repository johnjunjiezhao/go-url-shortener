package store

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// Define a struct to hold the Redis client
type StoreService struct {
	redisClient *redis.Client
}

// Top level declaration for the store service and context
var (
	storeService = &StoreService{}
	ctx          = context.Background()
)

const CacheDuration = time.Hour * 6 // Cache duration for the shortened URLs

// Initialize the Redis client
func InitializeStore() *StoreService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	println("Redis connected:", pong)

	storeService.redisClient = redisClient
	return storeService
}

func SaveURLMapping(shortCode string, originalURL string, userId string) {
	err := storeService.redisClient.Set(ctx, shortCode, originalURL, CacheDuration).Err()
	if err != nil {
		panic(err)
	}
}

func RetrieveOriginalURL(shortCode string) string {
	result, err := storeService.redisClient.Get(ctx, shortCode).Result()
	if err != nil {
		panic(err)
	}
	return result
}
