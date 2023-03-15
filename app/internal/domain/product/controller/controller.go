package controller

import (
	"errors"
	"fmt"
	"my_project/internal/domain/product/dto"
	"my_project/internal/domain/product/service"
	"my_project/internal/utils/customvalidator"
	"my_project/internal/utils/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductController interface {
	All(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type productController struct {
	service service.ProductService
}

func NewProductController(service service.ProductService) ProductController {
	return &productController{
		service: service,
	}
}

// GET ALL Roles
func (c *productController) All(ctx *fiber.Ctx) error {

	result := c.service.All()
	return ctx.Status(http.StatusOK).JSON(result)

}

//CREATE Role
func (c *productController) Create(ctx *fiber.Ctx) error {

	var productCreateDTO dto.ProductDTO
	err := ctx.BodyParser(&productCreateDTO)
	fmt.Println(productCreateDTO.Name)
	if err != nil {
		res := response.Error("Bad request", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = customvalidator.ValidateStruct(&productCreateDTO)
	if err != nil {
		res := response.Error("Validation error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = c.service.Create(&productCreateDTO)
	if err != nil {
		res := response.Error("Couldn't create", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(res)
	}

	res := response.Success(true, "Success", nil)
	return ctx.Status(http.StatusOK).JSON(res)

}

//UPDATE Role
func (c *productController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		res := response.Error("Error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	if id == 0 {
		res := response.Error("Error", "Invalid parameter", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	var productUpdateDTO dto.ProductDTO
	err = ctx.BodyParser(&productUpdateDTO)
	if err != nil {
		res := response.Error("Bad request", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = customvalidator.ValidateStruct(&productUpdateDTO)
	if err != nil {
		res := response.Error("Validation error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = c.service.Update(&productUpdateDTO, id)
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
func (c *productController) Delete(ctx *fiber.Ctx) error {
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
