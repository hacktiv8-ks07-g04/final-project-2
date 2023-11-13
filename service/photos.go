package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
)

type Photos interface{}

type PhotosImpl struct {
	repository repository.Photos
}

func NewPhotos(repository repository.Photos) *PhotosImpl {
	return &PhotosImpl{repository}
}
