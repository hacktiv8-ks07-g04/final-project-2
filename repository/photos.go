package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
)

type Photos interface {
	Add(photo *entity.Photo) (*entity.Photo, error)
	Get(photoId uint) (*entity.Photo, error)
	GetAll() ([]entity.Photo, error)
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

func (r *PhotosImpl) Get(photoId uint) (*entity.Photo, error) {
	var photo entity.Photo
	err := r.db.First(&photo, photoId).Error
	if err != nil {
		return nil, errors.New("photo not found")
	}

	return &photo, err
}

func (r *PhotosImpl) GetAll() ([]entity.Photo, error) {
	var photos []entity.Photo
	err := r.db.Find(&photos).Error
	if err != nil {
		return nil, errors.New("failed to get photos")
	}

	return photos, err
}
