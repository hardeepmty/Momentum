package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task represents a task item
type Task struct {
	ID          primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	UserID      primitive.ObjectID  `json:"userId" bson:"userId" validate:"required"`
	WorkspaceID *primitive.ObjectID `json:"workspaceId,omitempty" bson:"workspaceId,omitempty"`
	Title       string              `json:"title" bson:"title" validate:"required,min=1,max=200"`
	Description string              `json:"description" bson:"description" validate:"max=1000"`
	Tag         string              `json:"tag" bson:"tag" validate:"max=50"`
	Completed   bool                `json:"completed" bson:"completed"`
	Priority    string              `json:"priority" bson:"priority" validate:"oneof=low medium high"`
	DueDate     *time.Time          `json:"dueDate,omitempty" bson:"dueDate,omitempty"`
	Subtasks    []Subtask           `json:"subtasks" bson:"subtasks"`
	CreatedAt   time.Time           `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time           `json:"updatedAt" bson:"updatedAt"`
}