package repository

import (
	"context"
	"time"
	"todo-list-api/config"
	"todo-list-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TodoRepository defines data access methods for Todo items.
type TodoRepository interface {
	Create(todo *models.Todo) error
	Update(todo *models.Todo) error
	Delete(id primitive.ObjectID, userID primitive.ObjectID) error
	GetTodos(userID primitive.ObjectID, page, limit int64) ([]models.Todo, int64, error)
	GetByID(id primitive.ObjectID) (*models.Todo, error)
}

type todoRepository struct{}

// NewTodoRepository returns a new instance of TodoRepository.
func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (r *todoRepository) Create(todo *models.Todo) error {
	collection := config.DB.Collection("todos")
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	_, err := collection.InsertOne(context.Background(), todo)
	return err
}

func (r *todoRepository) Update(todo *models.Todo) error {
	collection := config.DB.Collection("todos")
	todo.UpdatedAt = time.Now()
	filter := bson.M{"_id": todo.ID, "user_id": todo.UserID}
	update := bson.M{"$set": bson.M{
		"title":       todo.Title,
		"description": todo.Description,
		"updated_at":  todo.UpdatedAt,
	}}
	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *todoRepository) Delete(id primitive.ObjectID, userID primitive.ObjectID) error {
	collection := config.DB.Collection("todos")
	filter := bson.M{"_id": id, "user_id": userID}
	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *todoRepository) GetTodos(userID primitive.ObjectID, page, limit int64) ([]models.Todo, int64, error) {
	collection := config.DB.Collection("todos")
	filter := bson.M{"user_id": userID}

	opts := options.Find()
	opts.SetSkip((page - 1) * limit)
	opts.SetLimit(limit)

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, 0, err
	}
	var todos []models.Todo
	if err := cursor.All(context.Background(), &todos); err != nil {
		return nil, 0, err
	}

	// Get total count
	total, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, 0, err
	}
	return todos, total, nil
}

func (r *todoRepository) GetByID(id primitive.ObjectID) (*models.Todo, error) {
	collection := config.DB.Collection("todos")
	var todo models.Todo
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}
