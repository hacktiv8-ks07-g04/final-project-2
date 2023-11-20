package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
)

type SocialMedias interface {
	Get(id uint) (*entity.SocialMedia, error)
	Add(socialMedia *entity.SocialMedia) (*entity.SocialMedia, error)
}

type SocialMediasImpl struct {
	db *gorm.DB
}

func NewSocialMedias(db *gorm.DB) *SocialMediasImpl {
	return &SocialMediasImpl{
		db: db,
	}
}

func (r *SocialMediasImpl) Get(id uint) (*entity.SocialMedia, error) {
	socialMedia := entity.SocialMedia{}

	if err := r.db.First(&socialMedia, id).Error; err != nil {
		return nil, errors.New("social media not found")
	}

	return &socialMedia, nil
}

func (r *SocialMediasImpl) Add(socialMedia *entity.SocialMedia) (*entity.SocialMedia, error) {
	if err := r.db.Create(&socialMedia).Error; err != nil {
		return nil, err
	}

	return socialMedia, nil
}
