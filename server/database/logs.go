package database

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	ID        string             `json:"id,omitempty" bson:"_id,omitempty"`
	ProjectID primitive.ObjectID `json:"project_id,omitempty" bson:"project_id,omitempty"`
	Message   string             `json:"message"`
	Type      string             `json:"type"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

func GetAllLogsForProjectById(c *fiber.Ctx, id string) ([]Log, error) {
	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	// get all records as a cursor
	query := bson.D{{Key: "project_id", Value: _id}, {Key: "type", Value: "error"}}

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

func CreateLogForProject(c *fiber.Ctx, id string) (Log, error) {
	// check that the given project id exists
	_, err := GetProjectById(c, id)

	if err != nil {
		return Log{}, err
	}

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return Log{}, err
	}
	collection := Ref.Db.Collection("logs")

	// New Log struct
	log := new(Log)
	// Parse body into struct
	if err := c.BodyParser(log); err != nil {
		return Log{}, err
	}

	// force MongoDB to always set its own generated ObjectIDs
	log.ID = ""
	log.ProjectID = _id
	log.CreatedAt = time.Now()
	log.UpdatedAt = time.Now()

	// insert the record
	insertionResult, err := collection.InsertOne(c.Context(), log)
	if err != nil {
		return Log{}, err
	}

	// get the just inserted record in order to return it as response
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	// decode the Mongo record into Log
	createdLog := &Log{}
	createdRecord.Decode(createdLog)

	// return the created Log in JSON format
	return *createdLog, nil
}
