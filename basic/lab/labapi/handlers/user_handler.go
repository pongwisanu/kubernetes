package handlers

import (
	"labapi/services"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userSrv services.UserService
}

func NewUserHandler(userSrv services.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userSrv.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(users)
}

func (h userHandler) GetUser(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	users, err := h.userSrv.GetUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(users)
}

func (h userHandler) AddUser(c *fiber.Ctx) error {

	request := services.UserRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).SendString(err.Error())
	}
	response, err := h.userSrv.AddUser(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(response)
}
