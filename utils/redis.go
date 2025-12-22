package utils

import (
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func NewRedisClient() error {

	redis_opt, err := redis.ParseURL(GetConfig().Redis.Host)
	if err != nil {
		return err
	}

	redis_opt.PoolSize = 10

	redisClient = redis.NewClient(redis_opt)
	return nil
}

func SetRedisClient(client *redis.Client) error {
	redisClient = client
	return nil
}

func GetRedisClient() *redis.Client {
	return redisClient
}
