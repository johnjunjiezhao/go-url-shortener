package store

import (
	"context"
	"os"
	"strconv"
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
	var (
		redisOpts *redis.Options
		err       error
	)

	if redisURL := os.Getenv("REDIS_URL"); redisURL != "" {
		redisOpts, err = redis.ParseURL(redisURL)
		if err != nil {
			panic(err)
		}
	} else {
		addr := os.Getenv("REDIS_ADDR")
		if addr == "" {
			addr = "localhost:6379"
		}

		password := os.Getenv("REDIS_PASSWORD")

		db := 0
		if v := os.Getenv("REDIS_DB"); v != "" {
			if parsed, convErr := strconv.Atoi(v); convErr == nil {
				db = parsed
			}
		}

		redisOpts = &redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		}
	}

	redisClient := redis.NewClient(redisOpts)

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
