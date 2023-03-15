package controller

import (
	"my_project/internal/domain/comment/dto"
	"my_project/internal/domain/comment/service"
	"my_project/internal/utils/customvalidator"
	"my_project/internal/utils/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CommentController interface {
	Create(ctx *fiber.Ctx) error
}

type commentController struct {
	service service.CommentService
}

func NewCommentController(service service.CommentService) CommentController {
	return &commentController{
		service: service,
	}
}

//CREATE Role
func (c *commentController) Create(ctx *fiber.Ctx) error {

	var commentCreateDTO dto.CommentDTO
	err := ctx.BodyParser(&commentCreateDTO)
	if err != nil {
		res := response.Error("Bad request", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = customvalidator.ValidateStruct(&commentCreateDTO)
	if err != nil {
		res := response.Error("Validation error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = c.service.Create(&commentCreateDTO)
	if err != nil {
		res := response.Error("Couldn't create", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(res)
	}

	res := response.Success(true, "Success", nil)
	return ctx.Status(http.StatusOK).JSON(res)

}
