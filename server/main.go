package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/immatheus/meitrics/server/api"
)

func main() {
	app := fiber.New()

	api.SetupRoutes(app)

	app.Listen(":3000")
}
