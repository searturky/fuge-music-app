package core

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis(conf *ConfYaml) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPassword,
		DB:       conf.RedisDB,
	})
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}

func CloseRedis() {
	RedisClient.Close()
}
