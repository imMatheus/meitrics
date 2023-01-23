package api

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

type Project struct {
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name"`
}

func Connect() error {
	uri := "mongodb+srv://meitrics:lMKJ9kByO9W45o1m@cluster0.bisw8pt.mongodb.net/meitrics?retryWrites=true&w=majority"
	dbName := "meitrics"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func SetupRoutes(app *fiber.App) {
	// Connect to the database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

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
	cursor, err := mg.Db.Collection("projects").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var projects []Project = make([]Project, 0)

	// iterate the cursor and decode each item into a Project
	if err := cursor.All(c.Context(), &projects); err != nil {
		return c.Status(500).SendString(err.Error())

	}
	// return projects list in JSON format
	return c.JSON(projects)

}
