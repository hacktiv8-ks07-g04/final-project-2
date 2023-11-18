package repository

import (
	"errors"
	"log"

	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
)

type Comments interface {
	Add(userID uint, comment *entity.Comment) (*entity.Comment, error)
}

type CommentsImpl struct {
	db *gorm.DB
}

func NewComments(db *gorm.DB) *CommentsImpl {
	return &CommentsImpl{db}
}

func (r *CommentsImpl) Add(userID uint, comment *entity.Comment) (*entity.Comment, error) {
	var user entity.User
	var photo entity.Photo

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, userID).Error; err != nil {
			return errors.New("user not found")
		}

		if err := tx.First(&photo, comment.PhotoID).Error; err != nil {
			return errors.New("photo not found")
		}

		comment.UserID = user.ID
		comment.User = user // because of validation
		comment.PhotoID = photo.ID
		comment.Photo = photo // because of validation
		comment.Photo.User = user

		log.Print(comment.User)

		if err := tx.Create(&comment).Error; err != nil {
			return err
		}

		return nil
	})

	return comment, err
}
