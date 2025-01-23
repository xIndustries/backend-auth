package models

import (
	"time"
)

// User represents the schema for the user entity stored in the database.
type User struct {
	ID        string    `json:"id" db:"id"`                       // Primary key (UUID)
	Auth0ID   string    `json:"auth0_id" db:"auth0_id"`           // Auth0 unique identifier
	Email     string    `json:"email" db:"email"`                 // User's email address
	Username  *string   `json:"username,omitempty" db:"username"` // Optional username
	CreatedAt time.Time `json:"created_at" db:"created_at"`       // Timestamp when the user was created
}

// CreateUserInput represents the data required to create a new user.
type CreateUserInput struct {
	Auth0ID  string `json:"auth0_id" validate:"required"`    // Required Auth0 ID
	Email    string `json:"email" validate:"required,email"` // Required email, validated
	Username string `json:"username,omitempty"`              // Optional username
}

// UpdateUserInput represents the data that can be updated for a user.
type UpdateUserInput struct {
	Auth0ID  string `json:"auth0_id" validate:"required"`               // Required to identify the user
	Username string `json:"username" validate:"omitempty,min=3,max=50"` // Optional username validation
}
