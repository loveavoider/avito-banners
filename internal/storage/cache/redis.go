package cache

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedis() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	return rdb
}