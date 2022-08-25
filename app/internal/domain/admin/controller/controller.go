package controller

import (
	"my_project/internal/domain/admin/dto"
	"my_project/internal/domain/admin/service"
	"my_project/internal/model"
	"my_project/internal/utils/customvalidator"
	"my_project/internal/utils/functions"
	"my_project/internal/utils/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"

	jService "my_project/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

type AdminController interface {
	All(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error

	Logout(ctx *fiber.Ctx) error
	Refresh(ctx *fiber.Ctx) error
}

type adminController struct {
	service         service.AdminService
	jwtAdminService jService.JWTAdminService
}

func NewAdminController(service service.AdminService, jwtAdminService jService.JWTAdminService) AdminController {
	return &adminController{
		service:         service,
		jwtAdminService: jwtAdminService,
	}
}

func (c *adminController) All(ctx *fiber.Ctx) error {

	result := c.service.All()
	return ctx.Status(http.StatusOK).JSON(result)

}

//CREATE Admin user
func (c *adminController) Create(ctx *fiber.Ctx) error {
	var adminCreateDTO dto.AdminDTO
	err := ctx.BodyParser(&adminCreateDTO)
	if err != nil {
		res := response.Error("Bad request", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = customvalidator.ValidateStruct(&adminCreateDTO)
	if err != nil {
		res := response.Error("Validation error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	_, err = c.service.FindByUsername(adminCreateDTO.Username)
	if err == nil {
		res := response.Error("Bu ulanyjy ady öňem bar", "Bu ulanyjy ady öňem bar", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = c.service.Create(&adminCreateDTO)
	if err != nil {
		res := response.Error("Couldn't create", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(res)
	}

	res := response.Success(true, "Success", nil)
	return ctx.Status(http.StatusOK).JSON(res)
}

type response_login struct {
	TokenDetails *model.TokenDetails `json:"tokenDetails"`
	Userdata     *model.Admin        `json:"userdata"`
}

//Login Admin user
func (c *adminController) Login(ctx *fiber.Ctx) error {
	var adminLoginDTO dto.AdminLoginDTO
	err := ctx.BodyParser(&adminLoginDTO)
	if err != nil {
		res := response.Error("Bad request", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = customvalidator.ValidateStruct(&adminLoginDTO)
	if err != nil {
		res := response.Error("Validation error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	userdata, err := c.service.FindByUsername(adminLoginDTO.Username)
	if err != nil {
		res := response.Error("Bu ulanyjy yok", "Bu ulanyjy yok", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	comparedPassword := functions.ComparePassword(userdata.Password, []byte(adminLoginDTO.Password))
	if !comparedPassword {
		res := response.Error("Invalid credentials", "Invalid credentials", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	// user_id := strconv.FormatUint(uint64(userdata.ID), 10)
	// fmt.Println(reflect.TypeOf(s))

	// createJwtToken, err := jService.CreateToken(user_id)
	createJwtToken, err := c.jwtAdminService.CreateToken(strconv.FormatUint(uint64(userdata.ID), 10), userdata.Username)

	if err != nil {
		res := response.Error("Token doredilmedi", "Token doredilmedi", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	c.service.CreateAuth(strconv.FormatUint(userdata.ID, 10), createJwtToken)
	// 1
	// data := map[string]interface{}{
	// 	"userdata":       userdata,
	// 	"createJwtToken": createJwtToken,
	// }
	// 2
	// var response_data response_login

	// response_data.TokenDetails = createJwtToken
	// response_data.Userdata = userdata

	userdata.AccessToken = createJwtToken.AccessToken
	userdata.RefreshToken = createJwtToken.RefreshToken

	return ctx.Status(http.StatusOK).JSON(userdata)

}

//LOGOUT ADMIN
func (c *adminController) Logout(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	headerParts := strings.Split(authHeader, " ")
	token, err := c.jwtAdminService.ValidateAdminAccessToken(headerParts[1])
	if err != nil {
		res := response.Error("Error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			res := response.Error("Error", err.Error(), nil)
			return ctx.Status(http.StatusUnauthorized).JSON(res)
		}
		deleted, err := c.service.DeleteAuth(accessUuid)
		if err != nil || deleted == 0 {
			res := response.Error("Error", "errror", nil)
			return ctx.Status(http.StatusInternalServerError).JSON(res)
		}
		res := response.Success(true, "OK", nil)
		return ctx.Status(http.StatusOK).JSON(res)
	}
	res := response.Error("Invalid token", err.Error(), nil)
	return ctx.Status(http.StatusBadRequest).JSON(res)
}

//REFRESH TOKEN
func (c *adminController) Refresh(ctx *fiber.Ctx) error {
	var adminRefreshTokenDTO dto.RefreshTokenDTO
	err := ctx.BodyParser(&adminRefreshTokenDTO)
	if err != nil {
		res := response.Error("Error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = customvalidator.ValidateStruct(adminRefreshTokenDTO)
	if err != nil {
		res := response.Error("Validation error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	token, err := c.jwtAdminService.ValidateAdminRefreshToken(adminRefreshTokenDTO.RefreshToken)
	if err != nil {
		res := response.Error("Error", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(res)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string)
		if !ok {
			res := response.Error("Error", err.Error(), nil)
			return ctx.Status(http.StatusInternalServerError).JSON(res)
		}
		userId, ok := claims["user_id"].(string)
		if !ok {
			res := response.Error("Error", err.Error(), nil)
			return ctx.Status(http.StatusInternalServerError).JSON(res)
		}
		username, ok := claims["username"].(string)
		if !ok {
			res := response.Error("Error", err.Error(), nil)
			return ctx.Status(http.StatusInternalServerError).JSON(res)
		}
		deleted, err := c.service.DeleteAuth(refreshUuid)
		if err != nil || deleted == 0 {
			res := response.Error("Error", err.Error(), nil)
			return ctx.Status(http.StatusInternalServerError).JSON(res)
		}

		// createJwtToken, err := jService.CreateToken(user_id)

		generatedToken, err := c.jwtAdminService.CreateToken(userId, username)
		// user_id := strconv.FormatUint(uint64(userdata.ID), 10)
		// fmt.Println(reflect.TypeOf(s))

		// generatedToken, err := jService.CreateToken(user_id)

		if err != nil {
			res := response.Error("Error", err.Error(), nil)
			return ctx.Status(http.StatusInternalServerError).JSON(res)
		}
		err = c.service.CreateAuth(userId, generatedToken)
		if err != nil {
			res := response.Error("Error", err.Error(), nil)
			return ctx.Status(http.StatusInternalServerError).JSON(res)
		}
		user, err := c.service.FindByUsername(username)
		if err != nil {
			res := response.Error("Error", err.Error(), nil)
			return ctx.Status(http.StatusNotFound).JSON(res)
		}
		user.AccessToken = generatedToken.AccessToken
		user.RefreshToken = generatedToken.RefreshToken
		res := response.Success(true, "OK", user)
		return ctx.Status(http.StatusOK).JSON(res)
	}
	res := response.Error("Invalid token", err.Error(), nil)
	return ctx.Status(http.StatusBadRequest).JSON(res)
}

//UPDATE Admin user
func (c *adminController) Update(ctx *fiber.Ctx) error {
	var adminUpdateDTO dto.AdminDTO
	id, err := ctx.ParamsInt("id")
	if err != nil {
		res := response.Error("Error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	if id == 0 {
		res := response.Error("Error", "Invalid parameter", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	err = ctx.BodyParser(&adminUpdateDTO)
	if err != nil {
		res := response.Error("Bad request", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = customvalidator.ValidateStruct(&adminUpdateDTO)
	if err != nil {
		res := response.Error("Validation error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	err = c.service.Update(&adminUpdateDTO, id)
	if err != nil {
		res := response.Error("Couldn't update", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(res)
	}
	res := response.Success(true, "Success", nil)
	return ctx.Status(http.StatusOK).JSON(res)
}
