package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
)

type Users interface {
	Register(r *dto.RegisterRequest) (*entity.User, error)
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
