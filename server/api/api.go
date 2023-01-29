package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/immatheus/meitrics/server/database"
)

func SetupRoutes(app *fiber.App) {
	// Connect to the database
	if err := database.Ref.Connect(); err != nil {
		log.Fatal(err)
	}

	app.Use(cors.New())

	// Welcome endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/health", Health)
	app.Get("/projects", GetAllProjects)
	app.Post("/project", CreateProject)
	app.Get("/projects/:id", GetProjectById)
	app.Patch("/projects/:id/update-private-key", UpdatePrivateKeyForProjectById)
	app.Get("/projects/:id/logs", GetLogForProject)
	app.Post("/log", CreateLog)
}

func Health(c *fiber.Ctx) error {
	return c.SendString("Ok!")
}
