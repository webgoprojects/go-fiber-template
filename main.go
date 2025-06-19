package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/linesmerrill/go-fiber-template/routes"
)

func main() {
	app := fiber.New()

	routes.RegisterHealthRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
} 