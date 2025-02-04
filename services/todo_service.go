package services

import (
	"todo-list-api/models"
	"todo-list-api/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TodoService is the business logic layer for managing Todo items.
type TodoService interface {
	CreateTodo(todo *models.Todo) error
	UpdateTodo(todo *models.Todo) error
	DeleteTodo(id string, userID string) error
	GetTodos(userID string, page, limit int64) ([]models.Todo, int64, error)
	GetTodoByID(id string) (*models.Todo, error)
}

type todoService struct {
	todoRepo repository.TodoRepository
}

// NewTodoService returns a new instance of TodoService.
func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{todoRepo}
}

func (s *todoService) CreateTodo(todo *models.Todo) error {
	return s.todoRepo.Create(todo)
}

func (s *todoService) UpdateTodo(todo *models.Todo) error {
	return s.todoRepo.Update(todo)
}

func (s *todoService) DeleteTodo(id string, userID string) error {
	todoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	return s.todoRepo.Delete(todoID, userObjID)
}

func (s *todoService) GetTodos(userID string, page, limit int64) ([]models.Todo, int64, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, 0, err
	}
	return s.todoRepo.GetTodos(userObjID, page, limit)
}

func (s *todoService) GetTodoByID(id string) (*models.Todo, error) {
	todoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return s.todoRepo.GetByID(todoID)
}
