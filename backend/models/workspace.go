package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workspace struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID   `json:"userId" bson:"userId" validate:"required"`
	Name      string               `json:"name" bson:"name" validate:"required,min=1,max=100"`
	Tasks     []primitive.ObjectID `json:"tasks" bson:"tasks"`
	CreatedAt time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time            `json:"updatedAt" bson:"updatedAt"`
}