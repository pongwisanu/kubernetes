package main

import (
	"labapi/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	routes.Register(app)

	app.Use(cors.New())

	app.Listen(":8000")
}
