package constructor

import (
	"my_project/internal/domain/permission/controller"
	"my_project/internal/domain/permission/service"
	"my_project/internal/domain/permission/storage"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	PermissionRepository storage.PermissionStorage
	PermissionService    service.PermissionService
	PermissionController controller.PermissionController
)

func PermissionRequirementsCreator(client *gorm.DB, redis *redis.Client) {
	PermissionRepository = storage.NewPermissionStorage(client)
	PermissionService = service.NewPermissionService(PermissionRepository, redis)
	PermissionController = controller.NewPermissionController(PermissionService)
}
