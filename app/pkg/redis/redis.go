package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type redisConfig struct {
	dsn string
}

// NewPgConfig creates new pg config instance
func NewRedisConfig(dsn string) *redisConfig {
	return &redisConfig{
		dsn: dsn,
	}
}

func NewRedisClient(cfg *redisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.dsn, //redis port
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to redis...")
	}
	return client
}

func CloseRedis(client *redis.Client) {
	err := client.Close()
	if err != nil {
		panic("Failed to close connection from redis")
	}
}
