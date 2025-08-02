package cache

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	Client *redis.Client
	Ctx    = context.Background()
)

// InitRedis initializes the Redis client using environment variables.
func InitRedis() {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("REDIS_PORT")
	if port == "" {
		port = "6379"
	}

	password := os.Getenv("REDIS_PASSWORD") // Optional

	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})

	// Use a timeout context to avoid hanging on ping
	ctx, cancel := context.WithTimeout(Ctx, 3*time.Second)
	defer cancel()

	if _, err := Client.Ping(ctx).Result(); err != nil {
		panic("❌ Redis connection failed: " + err.Error())
	}

	fmt.Println("✅ Redis connected successfully")
}

// SetCache sets a value in Redis with the given TTL.
func SetCache(key string, value string, ttl time.Duration) error {
	ctx, cancel := context.WithTimeout(Ctx, 2*time.Second)
	defer cancel()

	return Client.Set(ctx, key, value, ttl).Err()
}

// GetCache retrieves a value by key. Returns empty string and nil error if key not found.
func GetCache(key string) (string, error) {
	ctx, cancel := context.WithTimeout(Ctx, 2*time.Second)
	defer cancel()

	result, err := Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		// Key does not exist
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return result, nil
}

// DeleteCache deletes a key from Redis.
func DeleteCache(key string) error {
	ctx, cancel := context.WithTimeout(Ctx, 2*time.Second)
	defer cancel()

	return Client.Del(ctx, key).Err()
}
