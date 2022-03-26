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

	redisPort := os.Getenv("REDIS_PORT")
	redisPass := os.Getenv("REDIS_PASS")
	redisHost := os.Getenv("REDIS_HOST")

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPass,
		DB:       0, // use default DB
	})

	return rdb
}
