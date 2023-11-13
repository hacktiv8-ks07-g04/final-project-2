package repository

import (
	"gorm.io/gorm"
)

type Users interface{}

type UsersImpl struct {
	db *gorm.DB
}

func NewUsers(db *gorm.DB) *UsersImpl {
	return &UsersImpl{db}
}
