package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Todo represents a task or to-do list item.
type Todo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
