package controller

import (
	"online-articles/common"
	"online-articles/configuration"
	"online-articles/exception"
	"online-articles/model"
	"online-articles/service"

	"github.com/gofiber/fiber/v2"
)

func NewUserController(userService *service.UserService, config configuration.Config) *UserController {
	return &UserController{UserService: *userService, Config: config}
}

type UserController struct {
	service.UserService
	configuration.Config
}

func (controller UserController) Route(app *fiber.App) {
	app.Post("/v1/api/login", controller.LoginAuthentication)
	app.Post("/v1/api/register", controller.RegisterUser)
}

// LoginAuthentication func LoginAuthentication.
// @Description LoginAuthentication.
// @Summary LoginAuthentication
// @Tags User
// @Accept json
// @Produce json
// @Param request body model.UserModel true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/login [post]
func (controller UserController) LoginAuthentication(c *fiber.Ctx) error {
	var request model.UserModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	result := controller.UserService.LoginAuthentication(c.Context(), request)
	tokenJwtResult := common.GenerateToken(result.Username, controller.Config)
	resultWithToken := map[string]interface{}{
		"token":    tokenJwtResult,
		"username": result.Username,
	}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    resultWithToken,
	})
}

// RegisterUser func RegisterUser.
// @Description RegisterUser.
// @Summary RegisterUser
// @Tags User
// @Accept json
// @Produce json
// @Param request body model.UserModel true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/register [post]
func (controller UserController) RegisterUser(c *fiber.Ctx) error {
	var request model.UserModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	response := controller.UserService.RegisterUser(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}
