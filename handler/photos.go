package handler

import (
	"github.com/hacktiv8-ks07-g04/final-project-2/service"
)

type Photos interface{}

type PhotosImpl struct {
	service service.Photos
}

func NewPhotos(service service.Photos) *PhotosImpl {
	return &PhotosImpl{service}
}
