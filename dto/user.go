package dto

import (
	"time"
)

type User struct {
	ID        uint       `json:"id,omitempty"`
	Username  string     `json:"username"              binding:"required"`
	Email     string     `json:"email"                 binding:"required,email"`
	Password  string     `json:"password,omitempty"    binding:"required,min=6"`
	Age       uint       `json:"age"                   binding:"required"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Age      uint   `json:"age"      binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"    binding:"email"`
}

type UserResponse struct {
	ID        uint       `json:"id,omitempty"`
	Username  string     `json:"username"`
	Email     string     `json:"email,omitempty"`
	Age       uint       `json:"age,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
