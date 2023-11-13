package handler

import (
	"github.com/hacktiv8-ks07-g04/final-project-2/service"
)

type Users interface{}

type UsersImpl struct {
	service service.Users
}

func NewUsers(service service.Users) *UsersImpl {
	return &UsersImpl{service}
}
