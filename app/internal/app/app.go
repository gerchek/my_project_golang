package app

import (
	"time"

	"my_project/internal/setup/routes"
	"my_project/internal/utils/response"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func NewApp(logger *logrus.Logger) (app *fiber.App) {

	app = fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			logger.WithFields(logrus.Fields{
				"baseUrl": ctx.Hostname() + ctx.OriginalURL(),
				"user":    ctx.Get("user"),
			}).Error(err.Error())
			res := response.Error("Fiber error handler", err.Error(), nil)
			result := ctx.Status(code).JSON(res)
			return result
		},

		AppName:                 "My Project",
		BodyLimit:               4 * 1024 * 1024,
		EnableTrustedProxyCheck: true,
		ServerHeader:            "Döwlet hyzmatlar portaly",
		WriteTimeout:            time.Second * 40,
		ReadTimeout:             time.Second * 40,
	})

	// fmt.Println(redisClient)

	routes.SetAllAdminRoutes(app)

	return app
}
