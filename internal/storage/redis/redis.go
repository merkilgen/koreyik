package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/serwennn/koreyik/internal/config"
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
