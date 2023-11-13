package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
)

type SocialMedias interface{}

type SocialMediasImpl struct {
	repository repository.SocialMedias
}

func NewSocialMedias(repository repository.SocialMedias) *SocialMediasImpl {
	return &SocialMediasImpl{repository}
}
