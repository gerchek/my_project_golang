package routes

import (
	adminConstructor "my_project/internal/domain/admin/constructor"
	permissionConstructor "my_project/internal/domain/permission/constructor"
	roleConstructor "my_project/internal/domain/role/constructor"

	"my_project/internal/middleware"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func SetAllAdminRoutes(app *fiber.App, redisClient *redis.Client) {
	adminApi := app.Group("/api/", middleware.CheckAdminAPIToken)

	// auth routes
	adminAuth := adminApi.Group("/admin")
	adminAuth.Get("/all", adminConstructor.AdminController.All)
	adminAuth.Post("/register", adminConstructor.AdminController.Create)
	adminAuth.Post("/login", adminConstructor.AdminController.Login)
	adminAuth.Post("/refresh", adminConstructor.AdminController.Refresh)
	adminAuth.Post("/update/:id", middleware.IsAdminAuthenticate(redisClient), adminConstructor.AdminController.Update)

	adminAuth.Get("/logout", middleware.IsAdminAuthenticate(redisClient), adminConstructor.AdminController.Logout)

	// role routes
	role := adminApi.Group("/role")
	role.Get("/all", roleConstructor.RoleController.All)
	role.Post("/create", roleConstructor.RoleController.Create)
	role.Delete("/delete/:id", roleConstructor.RoleController.Delete)

	// permission routes
	permission := adminApi.Group("/permission")
	permission.Get("/all", permissionConstructor.PermissionController.All)
	permission.Post("/create", permissionConstructor.PermissionController.Create)

}
