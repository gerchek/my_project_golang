package constructor

import (
	"my_project/internal/domain/role/controller"
	"my_project/internal/domain/role/service"
	"my_project/internal/domain/role/storage"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	RoleRepository storage.RoleStorage
	RoleService    service.RoleService
	RoleController controller.RoleController
)

func RoleRequirementsCreator(client *gorm.DB, redis *redis.Client) {
	RoleRepository = storage.NewRoleStorage(client)
	RoleService = service.NewRoleService(RoleRepository, redis)
	RoleController = controller.NewRoleController(RoleService)
}
