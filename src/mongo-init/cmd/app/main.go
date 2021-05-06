package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading env variables: %s \n", err.Error())
	}
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%d/"+
		"?authSource=%s&readPreference=primary&ssl=false",
		getEnv("MONGO_USER", ""),
		getEnv("MONGO_PASS", ""),
		getEnv("MONGO_HOST", ""),
		getEnvAsInt("MONGO_PORT", 27017),
		getEnv("MONGO_DB", ""))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := NewMongoDB(mongoURI, ctx)
	if err != nil {
		log.Fatalf("error on trying to connect to the mongodb: %s", err.Error())
	}

	initDb(db, ctx)
}

func initDb(db *mongo.Client, ctx context.Context) {
	if dbName, exists := os.LookupEnv("MONGO_DB"); exists {
		if collectionName, exists := os.LookupEnv("MONGO_COLLECTION"); exists {
			var err = db.Database(dbName).
				CreateCollection(ctx, collectionName)
			if err != nil {
				log.Fatalf("error creating collection %s", err.Error())
			}
		}
		_ = db.Disconnect(ctx)
	}
}

func NewMongoDB(mongoURI string, ctx context.Context) (*mongo.Client, error) {
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
	return client, err
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
