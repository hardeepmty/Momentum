package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workspace struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID   `json:"userId" bson:"userId"`
	Name      string               `json:"name" bson:"name"`
	Tasks     []primitive.ObjectID `json:"tasks" bson:"tasks"` // References Task IDs
	CreatedAt time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time            `json:"updatedAt" bson:"updatedAt"`
}
