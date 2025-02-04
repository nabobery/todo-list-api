package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a registered user in the system.
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"` // omit in responses
}
