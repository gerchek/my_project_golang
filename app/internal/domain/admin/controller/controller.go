package controller

import (
	"fmt"
	"my_project/internal/domain/admin/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AdminController interface {
	Test(ctx *fiber.Ctx) error
}

type adminController struct {
	service service.AdminService
}

func NewAdminController(service service.AdminService) AdminController {
	fmt.Println("domain/admin/controller/NewAdminController")
	return &adminController{
		service: service,
	}
}

func (c *adminController) Test(ctx *fiber.Ctx) error {

	result := c.service.All()
	return ctx.Status(http.StatusOK).JSON(result)

}
