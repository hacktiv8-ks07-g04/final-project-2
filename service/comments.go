package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
)

type Comments interface{}

type CommentsImpl struct {
	repository repository.Comments
}

func NewComments(repository repository.Comments) *CommentsImpl {
	return &CommentsImpl{repository}
}
