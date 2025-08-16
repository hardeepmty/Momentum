package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Username  string               `json:"username" bson:"username"`
	Email     string               `json:"email" bson:"email"`
	Password  string               `json:"-" bson:"password"`
	Workspaces []primitive.ObjectID `json:"workspaces" bson:"workspaces"`
	CreatedAt time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time            `json:"updatedAt" bson:"updatedAt"`
}


type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRequest struct {
	AuthRequest
	Username string `json:"username"`
}