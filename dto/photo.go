package dto

import (
	"time"
)

type Photo struct {
	ID        uint       `json:"id,omitempty"`
	Title     string     `json:"title"                binding:"required"`
	Caption   string     `json:"caption"`
	PhotoURL  string     `json:"photo_url"            binding:"required"`
	UserID    uint       `json:"user_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type AddPhotoRequest struct {
	Title    string `json:"title"     binding:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" binding:"required"`
}

type UpdatePhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}

type AddPhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UserPhoto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type GetPhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      UserPhoto `json:"user"`
}

type PhotoResponse struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoURL  string     `json:"photo_url"`
	UserID    uint       `json:"user_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type DeletePhotoResponse struct {
	Message string `json:"message"`
}
