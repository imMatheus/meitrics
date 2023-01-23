package api

import (
	"fmt"
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
}

func Health(c *fiber.Ctx) error {
	return c.SendString("Ok!")
}

func GetAllProjects(c *fiber.Ctx) error {
	projects, err := database.GetAllProjects(c)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// return projects list in JSON format
	return c.JSON(projects)

}

func GetProjectById(c *fiber.Ctx) error {
	// get id by params
	id := c.Params("id")
	fmt.Println("made it to route")
	fmt.Println(id)

	project, err := database.GetProjectById(c, id)

	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	// return projects list in JSON format
	return c.JSON(project)

}
