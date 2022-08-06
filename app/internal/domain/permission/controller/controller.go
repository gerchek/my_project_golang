package controller

import (
	"my_project/internal/domain/permission/dto"
	"my_project/internal/domain/permission/service"
	"my_project/internal/utils/customvalidator"
	"my_project/internal/utils/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type PermissionController interface {
	All(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}

type permissionController struct {
	service service.PermissionService
}

func NewPermissionController(service service.PermissionService) PermissionController {
	return &permissionController{
		service: service,
	}
}

// GET ALL Permissions
func (c *permissionController) All(ctx *fiber.Ctx) error {

	result := c.service.All()
	return ctx.Status(http.StatusOK).JSON(result)

}

//CREATE Permission
func (c *permissionController) Create(ctx *fiber.Ctx) error {
	var permissionCreateDTO dto.PermissionDTO
	err := ctx.BodyParser(&permissionCreateDTO)
	if err != nil {
		res := response.Error("Bad request", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = customvalidator.ValidateStruct(&permissionCreateDTO)
	if err != nil {
		res := response.Error("Validation error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = c.service.Create(&permissionCreateDTO)
	if err != nil {
		res := response.Error("Couldn't create", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	res := response.Success(true, "Success", nil)
	return ctx.Status(http.StatusOK).JSON(res)
}
