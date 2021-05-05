package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tasksvc/task-rest-api/internal/model"
	"tasksvc/task-rest-api/internal/repository/database"
)

type AuthorizationRepository interface {
}

type TaskRepository interface {
	CreateTask(task model.Task) (primitive.ObjectID, error)
	GetAllTasks() ([]model.Task, error)
	GetTaskById(id primitive.ObjectID) (model.Task, error)
	DeleteTaskById(id primitive.ObjectID) (bool, error)
	UpdateTaskById(id primitive.ObjectID, input model.Task) (bool, error)
}

type Repository struct {
	AuthorizationRepository
	TaskRepository
}

func NewRepository(db *database.MongoDB) *Repository {
	return &Repository{
		TaskRepository: NewTaskRepositoryMongoDB(db),
	}
}
