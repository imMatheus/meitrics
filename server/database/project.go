package database

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name"`
}

func GetAllProjects(c *fiber.Ctx) ([]Project, error) {
	// get all records as a cursor
	query := bson.D{{}}
	cursor, err := Ref.Db.Collection("projects").Find(c.Context(), query)
	if err != nil {
		return nil, err
	}

	var projects []Project = make([]Project, 0)

	// iterate the cursor and decode each item into a Project
	if err := cursor.All(c.Context(), &projects); err != nil {
		return nil, err
	}

	// return projects list in JSON format
	return projects, nil
}

func GetProjectById(c *fiber.Ctx, id string) (Project, error) {
	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return Project{}, err
	}

	query := bson.D{{Key: "_id", Value: _id}}

	var project Project
	if err := Ref.Db.Collection("projects").FindOne(c.Context(), query).Decode(&project); err != nil {
		return Project{}, err
	}

	return project, nil
}
