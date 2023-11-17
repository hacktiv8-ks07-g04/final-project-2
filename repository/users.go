package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
)

type Users interface {
	Register(u *entity.User) (*entity.User, error)
	Login(email string) (*entity.User, error)
	Update(id uint, data *dto.UpdateUserRequest) (*entity.User, error)
	Delete(id uint) error
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

func (r *UsersImpl) Login(email string) (*entity.User, error) {
	var u *entity.User

	err := r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("email = ?", email).First(&u).Error
		if err != nil {
			return errors.New("user not found")
		}

		return nil
	})

	return u, err
}

func (r *UsersImpl) Update(id uint, data *dto.UpdateUserRequest) (*entity.User, error) {
	var u *entity.User

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.User{}).Where("id = ?", id).First(&u).Error; err != nil {
			return err
		}

		if err := tx.Model(&u).Updates(data).Error; err != nil {
			return err
		}

		return nil
	})

	return u, err
}

func (r *UsersImpl) Delete(id uint) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&entity.User{}, id).Error; err != nil {
			return errors.New("user not found")
		}

		return nil
	})

	return err
}
