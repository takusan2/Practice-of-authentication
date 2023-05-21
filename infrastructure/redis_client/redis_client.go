package redis_client

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	SetSession(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	GetSession(ctx context.Context, key string) (string, error)
}

type redisClient struct {
	client *redis.Client
}

func NewRedisClient(client *redis.Client) RedisClient {
	return &redisClient{client: client}
}

func (s *redisClient) SetSession(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	fmt.Print("redisにセッションをセットします")
	return s.client.Set(ctx, key, value, expiration).Err()
}

func (s *redisClient) GetSession(ctx context.Context, key string) (string, error) {
	return s.client.Get(ctx, key).Result()
}
