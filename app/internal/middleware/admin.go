package middleware

import (
	"fmt"
	"my_project/internal/utils/response"
	"net/http"
	"strings"

	"my_project/internal/domain/admin/constructor"

	"my_project/internal/utils/functions"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
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

func IsAdminAuthenticate(redis *redis.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		// fmt.Println(authHeader)
		if len(authHeader) == 0 || authHeader == "" {
			res := response.Error("Error", "Invalid token or token not found 1", nil)
			return ctx.Status(http.StatusUnauthorized).JSON(res)
		}
		// fmt.Println(authHeader)
		headerParts := strings.Split(authHeader, " ")
		fmt.Println(headerParts[0])
		if len(headerParts) != 2 {
			res := response.Error("Error", "Invalid token or token not found 2", nil)
			return ctx.Status(http.StatusUnauthorized).JSON(res)
		}
		if headerParts[0] != "Bearer" {
			res := response.Error("Error", "Invalid token or token not found 3", nil)
			return ctx.Status(http.StatusUnauthorized).JSON(res)
		}
		token, err := constructor.JwtAdminService.ValidateAdminAccessToken(headerParts[1])
		if err != nil {
			res := response.Error("Error", "Invalid token or token not found 4", nil)
			return ctx.Status(http.StatusUnauthorized).JSON(res)
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			accessUuid, ok := claims["access_uuid"].(string)

			if !ok {
				res := response.Error("Error", "Invalid token or token not found 5", nil)
				return ctx.Status(http.StatusUnauthorized).JSON(res)
			}

			// userId, ok := claims["user_id"].(string)

			if !ok {
				res := response.Error("Error", "Invalid token or token not found 6", nil)
				return ctx.Status(http.StatusUnauthorized).JSON(res)
			}
			err = functions.IsAvailableOnRedis(accessUuid, redis)
			if err != nil {
				res := response.Error("Error", "Invalid token or token not found 7", nil)
				return ctx.Status(http.StatusUnauthorized).JSON(res)
			}
			return ctx.Next()
		} else {
			res := response.Error("Error", "Invalid token or token not found 8", nil)
			return ctx.Status(http.StatusUnauthorized).JSON(res)
		}
	}
}
