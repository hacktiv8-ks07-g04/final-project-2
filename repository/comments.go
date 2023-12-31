package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-2/dto"
)

type Comments interface {
	Create(comment *entity.Comment) (*entity.Comment, error)
	Get(id uint) (*entity.Comment, error)
	GetAll() ([]entity.Comment, error)
	Update(comment *entity.Comment, updatedData *dto.Comment) (*entity.Comment, error)
	Delete(comment *entity.Comment) error
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
		return nil, errors.New("comment not found")
	}

	return &comment, nil
}

func (r *CommentsImpl) Create(comment *entity.Comment) (*entity.Comment, error) {
	if err := r.db.Create(&comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *CommentsImpl) Update(comment *entity.Comment, updatedData *dto.Comment) (*entity.Comment, error) {
	if err := r.db.Model(&comment).Updates(&updatedData).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *CommentsImpl) GetAll() ([]entity.Comment, error) {
	comments := []entity.Comment{}

	if err := r.db.Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentsImpl) Delete(comment *entity.Comment) error {
	if err := r.db.Delete(&comment).Error; err != nil {
		return err
	}

	return nil
}
