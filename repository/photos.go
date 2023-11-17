package repository

import (
	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
)

type Photos interface {
	Add(userID uint, p *entity.Photo) (*entity.Photo, error)
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
