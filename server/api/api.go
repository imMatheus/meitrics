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
		return c.SendString("Hello, World 69!")
	})

	app.Get("/health", Health)
	app.Get("/projects", GetAllProjects)
	app.Get("/projects/:id", GetProjectById)
	app.Get("/projects/:id/logs", GetLogsForProject)
	app.Post("/projects/:id/log", CreateLog)
}

func Health(c *fiber.Ctx) error {
	return c.SendString("Ok!")
}
