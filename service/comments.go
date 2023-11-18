package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
)

type Comments interface {
	Add(userID uint, r *dto.CreateCommentRequest) (*entity.Comment, error)
}

type CommentsImpl struct {
	repository repository.Comments
}

func NewComments(repository repository.Comments) *CommentsImpl {
	return &CommentsImpl{repository}
}

func (s *CommentsImpl) Add(userID uint, r *dto.CreateCommentRequest) (*entity.Comment, error) {
	comment := &entity.Comment{
		Message: r.Message,
		UserID:  userID,
		PhotoID: r.PhotoID,
	}

	comment, err := s.repository.Add(userID, comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}
