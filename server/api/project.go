package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/immatheus/meitrics/server/database"
)

// https://github.com/gofiber/recipes/blob/master/mongodb/main.go

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

type CreateProjectDTO struct {
	Name string `json:"name"`
}

func CreateProject(c *fiber.Ctx) error {
	fmt.Println("made it to project route")

	project, err := database.CreateProject(c)

	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	// return projects list in JSON format
	return c.JSON(project)

}
