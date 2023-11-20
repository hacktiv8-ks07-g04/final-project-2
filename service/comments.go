package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-2/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
)

type Comments interface {
	Add(payload *dto.Comment) (*entity.Comment, error)
	GetAll() ([]entity.Comment, error)
	Update(comment *entity.Comment) (*entity.Comment, error)
}

type CommentsImpl struct {
	repository repository.Comments
}

func NewComments(repository repository.Comments) *CommentsImpl {
	return &CommentsImpl{repository}
}

func (s *CommentsImpl) Add(payload *dto.Comment) (*entity.Comment, error) {
	comment := entity.Comment{
		Message: payload.Message,
		PhotoID: payload.PhotoID,
		UserID:  payload.UserID,
	}

	return s.repository.Create(&comment)
}

func (s *CommentsImpl) GetAll() ([]entity.Comment, error) {
	return s.repository.GetAll()
}

func (s *CommentsImpl) Update(comment *entity.Comment) (*entity.Comment, error) {
	return s.repository.Update(comment)
}
