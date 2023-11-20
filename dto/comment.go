package dto

import (
	"time"
)

type Comment struct {
	ID        uint       `json:"id,omitempty"`
	Message   string     `json:"message"              binding:"required"`
	PhotoID   uint       `json:"photo_id"             binding:"required"`
	UserID    uint       `json:"user_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type AddCommentRequest struct {
	Message string `json:"message"  binding:"required"`
	PhotoID uint   `json:"photo_id" binding:"required"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" binding:"required"`
}

type CommentResponse struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	PhotoID   uint       `json:"photo_id"`
	UserID    uint       `json:"user_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
