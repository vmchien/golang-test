package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"zoro/pkg/config"
)

func GetInstance(config config.Config) *redis.Client {
	redisDB, _ := strconv.Atoi(config.RedisDB)
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.RedisHost, config.RedisPort),
		Password: "",
		DB:       redisDB,
	})
	client.Type(config.RedisType)
	client.TTL(config.RedisTTL)
	return client
}
