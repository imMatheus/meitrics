package api

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	// Welcome endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 69!")
	})

	app.Get("/health", Health)

}

func Health(c *fiber.Ctx) error {
	return c.SendString("Ok!")
}
