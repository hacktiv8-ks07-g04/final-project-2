package repository

import (
	"gorm.io/gorm"
)

type Comments interface{}

type CommentsImpl struct {
	db *gorm.DB
}

func NewComments(db *gorm.DB) *CommentsImpl {
	return &CommentsImpl{db}
}
