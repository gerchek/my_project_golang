package constructor

import (
	"fmt"

	adminConstructor "my_project/internal/domain/admin/constructor"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetConstructor(client *gorm.DB, redisClient *redis.Client, logger *logrus.Logger) {
	adminConstructor.AdminRequirementsCreator(client, redisClient)
	fmt.Println("SetConstructor")
}
