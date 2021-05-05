package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tasksvc/task-rest-api/internal/model"
	"tasksvc/task-rest-api/internal/repository"
)

type Authorization interface {
}

type Task interface {
	CreateTask(task model.Task) (primitive.ObjectID, error)
	GetAllTasks() ([]model.Task, error)
	GetTaskById(id primitive.ObjectID) (model.Task, error)
	DeleteTaskById(id primitive.ObjectID) (bool, error)
	UpdateTaskById(id primitive.ObjectID, input model.Task) (bool, error)
}

type Service struct {
	Authorization
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Task: NewTaskService(repos.TaskRepository),
	}
}
