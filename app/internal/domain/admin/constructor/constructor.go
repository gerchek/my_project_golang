package constructor

import (
	"my_project/internal/domain/admin/controller"
	"my_project/internal/domain/admin/service"
	"my_project/internal/domain/admin/storage"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	AdminRepository storage.AdminStorage
	adminService    service.AdminService
	AdminController controller.AdminController
)

func AdminRequirementsCreator(client *gorm.DB, redis *redis.Client) {
	AdminRepository = storage.NewAdminStorage(client)
	adminService = service.NewAdminService(AdminRepository, redis)
	AdminController = controller.NewAdminController(adminService)
}
