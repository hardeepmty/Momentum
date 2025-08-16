package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subtask struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Completed bool               `json:"completed" bson:"completed"`
}
