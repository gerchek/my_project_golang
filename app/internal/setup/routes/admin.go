package routes

import (
	adminConstructor "my_project/internal/domain/admin/constructor"
	"my_project/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetAllAdminRoutes(app *fiber.App) {
	adminApi := app.Group("/admin", middleware.CheckAdminAPIToken)
	adminApi.Get("/login", adminConstructor.Test)
}
