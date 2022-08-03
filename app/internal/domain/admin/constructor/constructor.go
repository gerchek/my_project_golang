package constructor

import (
	"github.com/gofiber/fiber/v2"
)

func Test(ctx *fiber.Ctx) error {
	// fmt.Println("constructor")
	return ctx.SendString("root")
}
