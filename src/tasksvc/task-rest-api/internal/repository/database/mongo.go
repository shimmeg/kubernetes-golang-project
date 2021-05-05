package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"tasksvc/task-rest-api/config"
	"time"
)

const (
	TasksDatabase   = "tasks-tracker"
	TasksCollection = "tasks"
)

type MongoDB struct {
	Client *mongo.Client
}

func (m *MongoDB) Open() {
	panic("implement me")
}

func (m *MongoDB) Close() {
	err := m.Client.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("Error on closing db: %s", err)
	}
}

func (m *MongoDB) GetCollection(name string) *mongo.Collection {
	collection := m.Client.Database(TasksDatabase).Collection(name)
	return collection
}

func NewMongoDB(config config.DBConfig) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%d/"+
		"?authSource=%s&readPreference=primary&ssl=false",
		config.Username, config.Password, config.Host, config.Port, config.DBName)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	databases, err := client.ListDatabases(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	log.Println(databases)
	return &MongoDB{client}, err
}
