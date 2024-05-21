package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetIndex(c *fiber.Ctx) error {
	return c.SendString("This is api for lab (url) but it hexagonal xD")
}

