package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetIndex(c *fiber.Ctx) error {

	x := `This is api for kubernetes lab
	/users 
	/user/:id
	/user/add
	`
	return c.SendString(x)
}
