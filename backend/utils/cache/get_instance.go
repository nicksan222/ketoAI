package cache

import (
	"sync"

	"github.com/nicksan222/ketoai/config"
	"github.com/redis/go-redis/v9"
)

var (
	clientInstance *redis.Client
	clientOnce     sync.Once
)

func GetCacheClient() (*Client, error) {
	clientOnce.Do(func() {
		env, err := config.LoadConfig()

		if err != nil {
			panic(err)
		}

		clientInstance = redis.NewClient(&redis.Options{
			Addr:     env.REDIS_HOST + ":" + env.REDIS_PORT,
			Password: env.REDIS_PASSWORD,
		})
	})

	return (*Client)(clientInstance), nil
}
