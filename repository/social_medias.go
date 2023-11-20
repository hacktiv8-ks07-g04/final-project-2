package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
)

type SocialMedias interface {
	Get(id uint) (*entity.SocialMedia, error)
	Add(socialMedia *entity.SocialMedia) (*entity.SocialMedia, error)
	GetAll() ([]entity.SocialMedia, error)
	Update(socialMedia *entity.SocialMedia) (*entity.SocialMedia, error)
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

func (r *SocialMediasImpl) GetAll() ([]entity.SocialMedia, error) {
	socialMedias := []entity.SocialMedia{}

	if err := r.db.Find(&socialMedias).Error; err != nil {
		return nil, err
	}

	return socialMedias, nil
}

func (r *SocialMediasImpl) Update(socialMedia *entity.SocialMedia) (*entity.SocialMedia, error) {
	if err := r.db.Save(&socialMedia).Error; err != nil {
		return nil, err
	}

	return socialMedia, nil
}
