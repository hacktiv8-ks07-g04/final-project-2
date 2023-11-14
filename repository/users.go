package repository

import (
	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
)

type Users interface {
	Register(u *entity.User) (*entity.User, error)
}

type UsersImpl struct {
	db *gorm.DB
}

func NewUsers(db *gorm.DB) *UsersImpl {
	return &UsersImpl{db}
}

func (r *UsersImpl) Register(u *entity.User) (*entity.User, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&u).Error
		if err != nil {
			return err
		}

		return nil
	})

	return u, err
}
