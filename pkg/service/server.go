package services

import (
	"github.com/go-redis/redis"
	"zoro/pkg/config"
	"zoro/pkg/db"
)

type MyServer struct {
	H   db.Handler
	Cof config.Config
	R   *redis.Client
}
