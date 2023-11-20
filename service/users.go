package service

import (
	"errors"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
	"github.com/hacktiv8-ks07-g04/final-project-2/utils"
)

type Users interface {
	Register(r *dto.RegisterRequest) (*entity.User, error)
	Login(r *dto.LoginRequest) (string, error)
	Update(id uint, r *dto.UpdateUserRequest) (*entity.User, error)
	Delete(id uint) error
	Get(id uint) (*entity.User, error)
}

type UsersImpl struct {
	repository repository.Users
}

func NewUsers(repository repository.Users) *UsersImpl {
	return &UsersImpl{repository}
}

func (s *UsersImpl) Register(r *dto.RegisterRequest) (*entity.User, error) {
	user := entity.User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
		Age:      r.Age,
	}

	u, err := s.repository.Register(&user)
	if err != nil {
		return nil, err
	}

	return u, err
}

func (s *UsersImpl) Login(r *dto.LoginRequest) (string, error) {
	user, err := s.repository.Login(r.Email)
	if err != nil {
		return "", errors.New("email or password is incorrect")
	}

	if err := utils.VerifyPassword(user.Password, r.Password); err != nil {
		return "", errors.New("email or password is incorrect")
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UsersImpl) Update(id uint, r *dto.UpdateUserRequest) (*entity.User, error) {
	user, err := s.repository.Update(id, r)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (s *UsersImpl) Delete(id uint) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *UsersImpl) Get(id uint) (*entity.User, error) {
	user, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
