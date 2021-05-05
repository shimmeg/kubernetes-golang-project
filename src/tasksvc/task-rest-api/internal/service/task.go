package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tasksvc/task-rest-api/internal/model"
	"tasksvc/task-rest-api/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task model.Task) (primitive.ObjectID, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]model.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) GetTaskById(id primitive.ObjectID) (model.Task, error) {
	return s.repo.GetTaskById(id)
}

func (s *TaskService) DeleteTaskById(id primitive.ObjectID) (bool, error) {
	return s.repo.DeleteTaskById(id)
}

func (s *TaskService) UpdateTaskById(id primitive.ObjectID, input model.Task) (bool, error) {
	return s.repo.UpdateTaskById(id, input)
}
