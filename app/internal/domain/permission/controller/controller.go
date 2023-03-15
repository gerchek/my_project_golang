package controller

import (
	"errors"
	"my_project/internal/domain/permission/dto"
	"my_project/internal/domain/permission/service"
	"my_project/internal/utils/customvalidator"
	"my_project/internal/utils/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PermissionController interface {
	All(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
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
	data, err := c.service.Create(&permissionCreateDTO)
	if err != nil {
		res := response.Error("Couldn't create", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	res := response.Success(true, "Success", data)
	return ctx.Status(http.StatusOK).JSON(res)
}

//UPDATE Permission
func (c *permissionController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		res := response.Error("Error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	if id == 0 {
		res := response.Error("Error", "Invalid parameter", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	var permissionUpdateDTO dto.PermissionDTO
	err = ctx.BodyParser(&permissionUpdateDTO)
	if err != nil {
		res := response.Error("Bad request", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = customvalidator.ValidateStruct(&permissionUpdateDTO)
	if err != nil {
		res := response.Error("Validation error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	data, err := c.service.Update(&permissionUpdateDTO, id)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			res := response.Error("Not found", err.Error(), nil)
			return ctx.Status(http.StatusNotFound).JSON(res)
		default:
			res := response.Error("Couldn't update", err.Error(), nil)
			return ctx.Status(http.StatusInternalServerError).JSON(res)
		}
	}
	res := response.Success(true, "Success", data)
	return ctx.Status(http.StatusOK).JSON(res)
}

//DELETE Permission
func (c *permissionController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		res := response.Error("Error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	if id == 0 {
		res := response.Error("Error", "Invalid parameter", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = c.service.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			res := response.Error("Not found", err.Error(), nil)
			return ctx.Status(http.StatusNotFound).JSON(res)
		default:
			res := response.Error("Couldn't delete", err.Error(), nil)
			return ctx.Status(http.StatusInternalServerError).JSON(res)
		}
	}
	res := response.Success(true, "Success", id)
	return ctx.Status(http.StatusOK).JSON(res)
}
