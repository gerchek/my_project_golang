package routes

import (
	adminConstructor "my_project/internal/domain/admin/constructor"
	commentsConstructor "my_project/internal/domain/comment/constructor"
	permissionConstructor "my_project/internal/domain/permission/constructor"
	productConstructor "my_project/internal/domain/product/constructor"
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
	role.Put("/update/:id", roleConstructor.RoleController.Update)

	// permission routes
	permission := adminApi.Group("/permission")
	permission.Get("/all", permissionConstructor.PermissionController.All)
	permission.Post("/create", permissionConstructor.PermissionController.Create)
	permission.Put("/update/:id", permissionConstructor.PermissionController.Update)
	permission.Delete("/delete/:id", permissionConstructor.PermissionController.Delete)

	// products routes
	products := adminApi.Group("/product")
	products.Get("/all", productConstructor.ProductController.All)
	products.Post("/create", productConstructor.ProductController.Create)
	products.Put("/update/:id", productConstructor.ProductController.Update)
	products.Delete("/delete/:id", productConstructor.ProductController.Delete)

	// comments routes
	comments := adminApi.Group("/comment")
	comments.Post("/create", commentsConstructor.CommentController.Create)

}
