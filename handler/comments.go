package handler

import (
	"github.com/hacktiv8-ks07-g04/final-project-2/service"
)

type Comments interface{}

type CommentsImpl struct {
	service service.Comments
}

func NewComments(service service.Comments) *CommentsImpl {
	return &CommentsImpl{service}
}
