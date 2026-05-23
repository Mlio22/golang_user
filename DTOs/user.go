package dtos

import "golang_user/models"

// UserInput used for create/update user
type UserInput struct {
	FirstName string `json:"first_name" binding:"required,min=1,max=100"`
	LastName  string `json:"last_name" binding:"required,min=1,max=100"`
	Email     string `json:"email" binding:"required,email,max=255"`
}

// UserResponse wraps a single user payload.
type UserResponse struct {
	Data models.User `json:"data"`
}

// UsersResponse wraps a list of users.
type UsersResponse struct {
	Data []models.User `json:"data"`
}
