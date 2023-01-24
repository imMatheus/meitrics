package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/immatheus/meitrics/server/database"
)

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
