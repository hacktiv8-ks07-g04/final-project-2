package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
)

type Users interface {
	Register(user *entity.User) (*entity.User, error)
	Login(email string) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(id uint) error
	Get(id uint) (*entity.User, error)
}

type UsersImpl struct {
	db *gorm.DB
}

func NewUsers(db *gorm.DB) *UsersImpl {
	return &UsersImpl{
		db: db,
	}
}

func (r *UsersImpl) Register(user *entity.User) (*entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UsersImpl) Login(email string) (*entity.User, error) {
	user := &entity.User{}

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *UsersImpl) Update(user *entity.User) (*entity.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UsersImpl) Delete(id uint) error {
	if err := r.db.Delete(&entity.User{}, id).Error; err != nil {
		return errors.New("user not found")
	}

	return nil
}

func (r *UsersImpl) Get(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, err
}
