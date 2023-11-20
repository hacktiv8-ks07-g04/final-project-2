package service

import (
	"errors"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-2/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
	"github.com/hacktiv8-ks07-g04/final-project-2/utils"
)

type Users interface {
	Register(payload *dto.RegisterRequest) (*entity.User, error)
	Login(payload *dto.LoginRequest) (string, error)
	Update(payload *dto.User) (*entity.User, error)
	Delete(id uint) error
	Get(id uint) (*entity.User, error)
}

type UsersImpl struct {
	userRepository repository.Users
}

func NewUsers(userRpository repository.Users) *UsersImpl {
	return &UsersImpl{
		userRepository: userRpository,
	}
}

func (s *UsersImpl) Register(payload *dto.RegisterRequest) (*entity.User, error) {
	user := entity.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
		Age:      payload.Age,
	}

	return s.userRepository.Register(&user)
}

func (s *UsersImpl) Login(payload *dto.LoginRequest) (string, error) {
	user, err := s.userRepository.Login(payload.Email)
	if err != nil {
		return "", errors.New("email or password is incorrect")
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		return "", errors.New("email or password is incorrect")
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UsersImpl) Update(payload *dto.User) (*entity.User, error) {
	user, err := s.userRepository.Get(payload.ID)
	if err != nil {
		return nil, err
	}

	user.Email = payload.Email
	user.Username = payload.Username

	return s.userRepository.Update(user)
}

func (s *UsersImpl) Delete(id uint) error {
	err := s.userRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *UsersImpl) Get(id uint) (*entity.User, error) {
	user, err := s.userRepository.Get(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
