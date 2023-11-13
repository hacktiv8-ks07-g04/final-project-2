package repository

import (
	"gorm.io/gorm"
)

type SocialMedias interface{}

type SocialMediasImpl struct {
	db *gorm.DB
}

func NewSocialMedias(db *gorm.DB) *SocialMediasImpl {
	return &SocialMediasImpl{db}
}
