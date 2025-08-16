package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Subtask represents a sub-task item
type Subtask struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title" validate:"required,min=1,max=200"`
	Completed bool               `json:"completed" bson:"completed"`
}