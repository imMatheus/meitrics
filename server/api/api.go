package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/immatheus/meitrics/server/database"
	"go.mongodb.org/mongo-driver/bson"
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
	app.Get("/projects", Projects)
}

func Health(c *fiber.Ctx) error {
	return c.SendString("Ok!")
}

func Projects(c *fiber.Ctx) error {
	// get all records as a cursor
	query := bson.D{{}}
	cursor, err := database.Ref.Db.Collection("projects").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var projects []database.Project = make([]database.Project, 0)

	// iterate the cursor and decode each item into a Project
	if err := cursor.All(c.Context(), &projects); err != nil {
		return c.Status(500).SendString(err.Error())

	}
	// return projects list in JSON format
	return c.JSON(projects)

}
