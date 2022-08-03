package middleware

import (
	"my_project/internal/utils/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CheckAdminAPIToken(ctx *fiber.Ctx) error {
	adminAPIToken := ctx.Get("API-KEY")
	if adminAPIToken != "E95E931F2187AF77157139B36137F" {
		res := response.Error("Permission denied", "API token not found", nil)
		return ctx.Status(http.StatusUnauthorized).JSON(res)
	}
	return ctx.Next()
}
