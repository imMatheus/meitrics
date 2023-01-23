package database

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"name"`
	ProjectID string    `json:"project_id,omitempty" bson:"project_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAllLogsForProjectById(c *fiber.Ctx, id string) ([]Log, error) {
	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	// get all records as a cursor
	query := bson.D{{Key: "_id", Value: _id}}

	cursor, err := Ref.Db.Collection("logs").Find(c.Context(), query)
	if err != nil {
		return nil, err
	}

	var logs []Log = make([]Log, 0)

	// iterate the cursor and decode each item into a Project
	if err := cursor.All(c.Context(), &logs); err != nil {
		return nil, err
	}

	// return projects list in JSON format
	return logs, nil
}
