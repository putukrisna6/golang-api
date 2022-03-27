package config

import (
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

func SetupRedisConnection() *redis.Client {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("failed to load env")
	}

	redisUrl := os.Getenv("REDIS_URL")

	opts, err := redis.ParseURL(redisUrl)
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opts)

	return rdb
}
