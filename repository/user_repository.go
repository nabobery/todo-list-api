package repository

import (
	"context"
	"todo-list-api/config"
	"todo-list-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRepository defines data access methods for User.
type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id primitive.ObjectID) (*models.User, error)
}

type userRepository struct{}

// NewUserRepository returns a new instance of UserRepository.
func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error {
	collection := config.DB.Collection("users")
	_, err := collection.InsertOne(context.Background(), user)
	return err
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	collection := config.DB.Collection("users")
	var user models.User
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id primitive.ObjectID) (*models.User, error) {
	collection := config.DB.Collection("users")
	var user models.User
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
