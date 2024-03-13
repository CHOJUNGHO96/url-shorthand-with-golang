package db

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(redisInfo map[string]string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisInfo["redisServer"], redisInfo["redisPort"]),
		Password: redisInfo["redisPassword"], // no password set
		DB:       0,                          // use default DB
	})
	return rdb
}
