package repository

import (
	"gorm.io/gorm"
)

type Photos interface{}

type PhotosImpl struct {
	db *gorm.DB
}

func NewPhotos(db *gorm.DB) *PhotosImpl {
	return &PhotosImpl{db}
}
