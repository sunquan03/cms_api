package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"time"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

func NewRedisClient() *redis.Client {
	addr := os.Getenv("redis_addr")
	password := os.Getenv("redis_password")

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	err := client.Ping(ctx).Err()
	if err != nil {
		log.Fatal(err)
	}

	return client
}
