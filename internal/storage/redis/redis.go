package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/serwennn/koreyik/internal/config"
)

var ctx = context.Background()

func New(cacheServerOptions config.CacheServer) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cacheServerOptions.Address,
		Password: cacheServerOptions.Password,
		DB:       cacheServerOptions.Database,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return client, nil
}
