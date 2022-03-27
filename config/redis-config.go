package config

import (
	"os"

	"github.com/go-redis/redis/v8"
)

func SetupRedisConnection() *redis.Client {
	redisUrl := os.Getenv("REDIS_URL")

	opts, err := redis.ParseURL(redisUrl)
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opts)

	return rdb
}
