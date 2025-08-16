package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the system
type User struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Username  string               `json:"username" bson:"username" validate:"required,min=3,max=50"`
	Email     string               `json:"email" bson:"email" validate:"required,email"`
	Password  string               `json:"-" bson:"password" validate:"required,min=8"`
	Workspaces []primitive.ObjectID `json:"workspaces" bson:"workspaces"`
	CreatedAt time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time            `json:"updatedAt" bson:"updatedAt"`
}

// AuthRequest represents authentication request data
type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// SignupRequest represents user registration data
type SignupRequest struct {
	AuthRequest
	Username string `json:"username" validate:"required,min=3,max=50"`
}