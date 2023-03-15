package constructor

import (
	adminConstructor "my_project/internal/domain/admin/constructor"
	commentConstructor "my_project/internal/domain/comment/constructor"
	permissionConstructor "my_project/internal/domain/permission/constructor"
	productConstructor "my_project/internal/domain/product/constructor"
	roleConstructor "my_project/internal/domain/role/constructor"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetConstructor(client *gorm.DB, redisClient *redis.Client, logger *logrus.Logger) {
	adminConstructor.AdminRequirementsCreator(client, redisClient)
	roleConstructor.RoleRequirementsCreator(client, redisClient)
	permissionConstructor.PermissionRequirementsCreator(client, redisClient)
	productConstructor.ProductRequirementsCreator(client, redisClient)
	commentConstructor.CommentRequirementsCreator(client, redisClient)
}
