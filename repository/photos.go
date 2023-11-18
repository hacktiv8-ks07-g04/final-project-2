package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
)

type Photos interface {
	Add(userID uint, p *entity.Photo) (*entity.Photo, error)
	GetAll() ([]entity.Photo, error)
	Update(photoID, userID uint, data *dto.AddPhotoRequest) (*entity.Photo, error)
	Delete(photoID, userID uint) error
}

type PhotosImpl struct {
	db *gorm.DB
}

func NewPhotos(db *gorm.DB) *PhotosImpl {
	return &PhotosImpl{db}
}

func (r *PhotosImpl) Add(userID uint, p *entity.Photo) (*entity.Photo, error) {
	var user entity.User

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, userID).Error; err != nil {
			return err
		}

		p.UserID = user.ID
		p.User = user // because of validation

		if err := tx.Create(&p).Error; err != nil {
			return err
		}

		return nil
	})

	return p, err
}

func (r *PhotosImpl) GetAll() ([]entity.Photo, error) {
	var photos []entity.Photo

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("User").Find(&photos).Error; err != nil {
			return err
		}

		return nil
	})

	return photos, err
}

func (r *PhotosImpl) Update(
	photoID, userID uint,
	data *dto.AddPhotoRequest,
) (*entity.Photo, error) {
	var photo entity.Photo

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.Photo{}).Preload("User").First(&photo, photoID).Error; err != nil {
			return errors.New("photo not found")
		}

		if photo.UserID != userID {
			return errors.New("you are not authorized to update this photo")
		}

		if err := tx.Model(&photo).Updates(data).Error; err != nil {
			return err
		}

		return nil
	})

	return &photo, err
}

func (r *PhotosImpl) Delete(photoID, userID uint) error {
	var photo entity.Photo

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.Photo{}).Preload("User").First(&photo, photoID).Error; err != nil {
			return errors.New("photo not found")
		}

		if photo.UserID != userID {
			return errors.New("you are not authorized to delete this photo")
		}

		if err := tx.Delete(&photo).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
