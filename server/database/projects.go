package database

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

func CreateProject(c *fiber.Ctx) (Project, error) {
	collection := Ref.Db.Collection("projects")

	// New Log struct
	project := new(Project)
	// Parse body into struct
	if err := c.BodyParser(project); err != nil {
		return Project{}, err
	}

	// validating that the client sent a valid message field
	if project.Name == "" {
		return Project{}, fmt.Errorf("please provide a valid name for your project")
	}

	// force MongoDB to always set its own generated ObjectIDs
	project.ID = ""
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()

	// insert the record
	insertionResult, err := collection.InsertOne(c.Context(), project)
	if err != nil {
		return Project{}, err
	}

	// get the just inserted record in order to return it as response
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	// decode the Mongo record into Log
	createdProject := &Project{}
	createdRecord.Decode(createdProject)

	// return the created Log in JSON format
	return *createdProject, nil
}
