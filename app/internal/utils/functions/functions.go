package functions

import (
	"context"

	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashedPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	return err == nil
}

func IsAvailableOnRedis(accessUuid string, redisClient *redis.Client) error {
	_, err := redisClient.Get(context.Background(), accessUuid).Result()
	if err != nil {
		return err
	}

	return nil
}
