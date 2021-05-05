package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tasksvc/task-rest-api/internal/model"
	"tasksvc/task-rest-api/internal/repository/database"
	"time"
)

type TaskMongoDB struct {
	db *database.MongoDB
}

func NewTaskRepositoryMongoDB(db *database.MongoDB) *TaskMongoDB {
	return &TaskMongoDB{db: db}
}

func (r *TaskMongoDB) CreateTask(task model.Task) (primitive.ObjectID, error) {
	collection := r.db.GetCollection(database.TasksCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, task)

	if err != nil {
		return primitive.ObjectID{}, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.ObjectID{}, errors.New("couldn't parse an inserted ID of the object")
	}
	return id, nil
}

func (r *TaskMongoDB) GetAllTasks() ([]model.Task, error) {
	collection := r.db.GetCollection(database.TasksCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var tasks []model.Task
	for cursor.Next(ctx) {
		var task model.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskMongoDB) GetTaskById(id primitive.ObjectID) (model.Task, error) {
	collection := r.db.GetCollection(database.TasksCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var task model.Task
	err := collection.FindOne(ctx, model.Task{Id: id}).Decode(&task)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (r *TaskMongoDB) DeleteTaskById(id primitive.ObjectID) (bool, error) {
	collection := r.db.GetCollection(database.TasksCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, model.Task{Id: id})
	if err != nil {
		return false, err
	}
	return result.DeletedCount == 1, nil
}

func (r *TaskMongoDB) UpdateTaskById(id primitive.ObjectID, task model.Task) (bool, error) {
	collection := r.db.GetCollection(database.TasksCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(
		ctx,
		model.Task{Id: id},
		bson.D{
			{"$set", bson.D{{"title", task.Title}, {"description", task.Description}}},
		},
	)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount == 1, nil
}
