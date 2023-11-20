package repository

import (
	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
)

type Photos interface {
	Add(photo *entity.Photo) (*entity.Photo, error)
}

type PhotosImpl struct {
	db *gorm.DB
}

func NewPhotos(db *gorm.DB) *PhotosImpl {
	return &PhotosImpl{db}
}

func (r *PhotosImpl) Add(photo *entity.Photo) (*entity.Photo, error) {
	err := r.db.Create(&photo).Error
	if err != nil {
		return nil, err
	}

	return photo, err
}
