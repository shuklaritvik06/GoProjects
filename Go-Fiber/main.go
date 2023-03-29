package main

import (
	"github.com/gofiber/fiber"
	"github.com/shuklaritvik06/GoProjects/fiber/config"
	"github.com/shuklaritvik06/GoProjects/fiber/routes"
)

func main() {
	config.Configure()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) {
		c.JSON(fiber.Map{"message": "Hello, World!"})
	})
	routes.AuthRoute(app)
	app.Listen(":3000")
}
