package red

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/serwennn/koreyik/internal/config"
	"time"
)

type CacheServer struct {
	client *redis.Client
}

var ctx = context.Background()

func New(cacheConfig config.CacheServer) (*CacheServer, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cacheConfig.Address,
		Password: cacheConfig.Password,
		DB:       cacheConfig.Database,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &CacheServer{client: client}, nil
}

func (c *CacheServer) Shutdown() error {
	return c.client.Close()
}

func (c *CacheServer) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *CacheServer) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

func (c *CacheServer) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

func (c *CacheServer) Keys(ctx context.Context, pattern string) ([]string, error) {
	return c.client.Keys(ctx, pattern).Result()
}
