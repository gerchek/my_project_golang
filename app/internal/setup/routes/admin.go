package routes

import (
	adminConstructor "my_project/internal/domain/admin/constructor"
	"my_project/internal/middleware"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func SetAllAdminRoutes(app *fiber.App, redisClient *redis.Client) {
	adminApi := app.Group("/admin", middleware.CheckAdminAPIToken)
	adminApi.Get("/all", adminConstructor.AdminController.All)
	adminApi.Post("/register", adminConstructor.AdminController.Create)
	adminApi.Post("/login", adminConstructor.AdminController.Login)
	adminApi.Post("/refresh", adminConstructor.AdminController.Refresh)

	adminApi.Get("/logout", middleware.IsAdminAuthenticate(redisClient), adminConstructor.AdminController.Logout)

}
