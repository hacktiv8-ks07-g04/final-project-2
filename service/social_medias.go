package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-2/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
)

type SocialMedias interface {
	Add(payload *dto.SocialMedia) (*entity.SocialMedia, error)
}

type SocialMediasImpl struct {
	repository repository.SocialMedias
}

func NewSocialMedias(repository repository.SocialMedias) *SocialMediasImpl {
	return &SocialMediasImpl{repository}
}

func (s *SocialMediasImpl) Add(payload *dto.SocialMedia) (*entity.SocialMedia, error) {
	socialMedia := entity.SocialMedia{
		Name:           payload.Name,
		SocialMediaURL: payload.SocialMediaURL,
		UserID:         payload.UserID,
	}

	return s.repository.Add(&socialMedia)
}
