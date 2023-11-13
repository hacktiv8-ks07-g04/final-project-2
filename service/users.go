package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
)

type Users interface{}

type UsersImpl struct {
	repository repository.Users
}

func NewUsers(repository repository.Users) *UsersImpl {
	return &UsersImpl{repository}
}
