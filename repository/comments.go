package repository

import (
	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
)

type Comments interface {
	Get(id uint) (*entity.Comment, error)
}

type CommentsImpl struct {
	db *gorm.DB
}

func NewComments(db *gorm.DB) *CommentsImpl {
	return &CommentsImpl{db}
}

func (r *CommentsImpl) Get(id uint) (*entity.Comment, error) {
	comment := entity.Comment{}

	if err := r.db.First(&comment, id).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}
