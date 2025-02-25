package initialize

import (
	"context"
	"fmt"
	"log"

	"github.com/csc13010-student-management/config"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func NewRedis(cfg config.RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println("Error initializing Redis")
		panic(err)
	}

	log.Println("Redis initialized successfully")
	return rdb
}
