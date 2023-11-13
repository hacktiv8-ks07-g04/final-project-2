package handler

import (
	"github.com/hacktiv8-ks07-g04/final-project-2/service"
)

type SocialMedias interface{}

type SocialMediasImpl struct {
	service service.SocialMedias
}

func NewSocialMedias(service service.SocialMedias) *SocialMediasImpl {
	return &SocialMediasImpl{service}
}
