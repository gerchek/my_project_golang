package constructor

import (
	"my_project/internal/domain/admin/controller"
	"my_project/internal/domain/admin/service"
	"my_project/internal/domain/admin/storage"
	jService "my_project/pkg/jwt"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	AdminRepository storage.AdminStorage
	adminService    service.AdminService
	JwtAdminService jService.JWTAdminService
	AdminController controller.AdminController
)

func AdminRequirementsCreator(client *gorm.DB, redis *redis.Client) {
	AdminRepository = storage.NewAdminStorage(client)
	adminService = service.NewAdminService(AdminRepository, redis)
	JwtAdminService = jService.NewJWTAdminService()
	AdminController = controller.NewAdminController(adminService, JwtAdminService)
}
