package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	UserID      primitive.ObjectID   `json:"userId" bson:"userId"`
	WorkspaceID *primitive.ObjectID  `json:"workspaceId,omitempty" bson:"workspaceId,omitempty"`
	Title       string               `json:"title" bson:"title"`
	Description string               `json:"description" bson:"description"`
	Tag         string               `json:"tag" bson:"tag"`
	Completed   bool                 `json:"completed" bson:"completed"`
	Priority    string               `json:"priority" bson:"priority"` // low, medium, high
	DueDate     *time.Time           `json:"dueDate,omitempty" bson:"dueDate,omitempty"`
	Subtasks    []Subtask            `json:"subtasks" bson:"subtasks"`
	CreatedAt   time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt" bson:"updatedAt"`
}
