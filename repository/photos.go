package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-2/dto"
)

type Photos interface {
	Add(photo *entity.Photo) (*entity.Photo, error)
	Get(photoId uint) (*entity.Photo, error)
	GetAll() ([]entity.Photo, error)
	Update(photoId uint, data *dto.UpdatePhotoRequest) (*entity.Photo, error)
	Delete(photoId uint) error
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

func (r *PhotosImpl) Update(
	photoId uint,
	data *dto.UpdatePhotoRequest,
) (*entity.Photo, error) {
	var photo *entity.Photo
	var err error

	photo, err = r.Get(photoId)
	if err != nil {
		return nil, err
	}

	if err := r.db.Model(&photo).Updates(data).Error; err != nil {
		return nil, err
	}

	return photo, err
}

func (r *PhotosImpl) Delete(photoId uint) error {
	err := r.db.Delete(&entity.Photo{}, photoId).Error
	if err != nil {
		return errors.New("failed to delete photo")
	}

	return nil
}
