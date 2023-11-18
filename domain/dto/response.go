package dto

import (
	"time"
)

type RegisterResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
}

type UpdateUserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Age       uint      `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AddPhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdatePhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserPhoto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type GetAllPhotosResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      UserPhoto `json:"user"`
}

type CreateCommentResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
