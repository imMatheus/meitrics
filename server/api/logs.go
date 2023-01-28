package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/immatheus/meitrics/server/database"
)

func GetLogForProject(c *fiber.Ctx) error {
	// get id by params
	id := c.Params("id")
	fmt.Println("made it to get logs route route")
	fmt.Println(id)

	logs, err := database.GetAllLogsForProjectById(c, id)

	fmt.Println("logs here")
	fmt.Println(logs)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	// return projects list in JSON format
	return c.JSON(logs)
}

func CreateLog(c *fiber.Ctx) error {
	fmt.Println("made it to lgo route")

	headers := c.GetReqHeaders()
	public := headers["Public-Key"]
	secret := headers["Secret-Key"]
	fmt.Println(headers)
	fmt.Println(public)
	fmt.Println(secret)

	id, err := database.ValidateKeysForProject(c, public, secret)

	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	project, err := database.CreateLogForProject(c, id)

	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	// return projects list in JSON format
	return c.JSON(project)

}
