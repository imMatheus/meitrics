package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Ref MongoInstance

type Project struct {
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name"`
}

func (m *MongoInstance) Connect() error {
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

	Ref = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}
