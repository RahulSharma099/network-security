package api

import (
	"github.com/RahulSharma099/network-security/api/controller"
	"github.com/gofiber/fiber/v2"
)


func handleRoot(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func setupApp(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", handleRoot)

	api.Get("/demo", controller.Demo)
}