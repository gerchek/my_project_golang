package controller

import (
	"errors"
	"my_project/internal/domain/role/dto"
	"my_project/internal/domain/role/service"
	"my_project/internal/utils/customvalidator"
	"my_project/internal/utils/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RoleController interface {
	All(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type roleController struct {
	service service.RoleService
}

func NewRoleController(service service.RoleService) RoleController {
	return &roleController{
		service: service,
	}
}

// GET ALL Roles
func (c *roleController) All(ctx *fiber.Ctx) error {

	result := c.service.All()
	return ctx.Status(http.StatusOK).JSON(result)

}

//CREATE Role
func (c *roleController) Create(ctx *fiber.Ctx) error {

	var roleCreateDTO dto.RoleDTO
	err := ctx.BodyParser(&roleCreateDTO)
	if err != nil {
		res := response.Error("Bad request", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = customvalidator.ValidateStruct(&roleCreateDTO)
	if err != nil {
		res := response.Error("Validation error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = c.service.Create(&roleCreateDTO)
	if err != nil {
		res := response.Error("Couldn't create", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(res)
	}

	res := response.Success(true, "Success", nil)
	return ctx.Status(http.StatusOK).JSON(res)

}

//UPDATE Role
func (c *roleController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		res := response.Error("Error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	if id == 0 {
		res := response.Error("Error", "Invalid parameter", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	var roleUpdateDTO dto.RoleDTO
	err = ctx.BodyParser(&roleUpdateDTO)
	if err != nil {
		res := response.Error("Bad request", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = customvalidator.ValidateStruct(&roleUpdateDTO)
	if err != nil {
		res := response.Error("Validation error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = c.service.Update(&roleUpdateDTO, id)
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

	res := response.Success(true, "Success", nil)
	return ctx.Status(http.StatusOK).JSON(res)
}

//DELETE ROLE
func (c *roleController) Delete(ctx *fiber.Ctx) error {
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
			res := response.Error("Couldn't update", err.Error(), nil)
			return ctx.Status(http.StatusInternalServerError).JSON(res)
		}
	}
	res := response.Success(true, "Success", nil)
	return ctx.Status(http.StatusOK).JSON(res)
}
