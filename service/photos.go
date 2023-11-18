package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
)

type Photos interface {
	Add(userID uint, r *dto.AddPhotoRequest) (*entity.Photo, error)
	GetAll() ([]entity.Photo, error)
	Update(photo *entity.Photo, r *dto.AddPhotoRequest) (*entity.Photo, error)
	Delete(photo *entity.Photo) error
}

type PhotosImpl struct {
	repository repository.Photos
}

func NewPhotos(repository repository.Photos) *PhotosImpl {
	return &PhotosImpl{repository}
}

func (s *PhotosImpl) Add(userID uint, r *dto.AddPhotoRequest) (*entity.Photo, error) {
	photo := entity.Photo{
		Title:    r.Title,
		Caption:  r.Caption,
		PhotoURL: r.PhotoURL,
	}

	p, err := s.repository.Add(userID, &photo)
	if err != nil {
		return nil, err
	}

	return p, err
}

func (s *PhotosImpl) GetAll() ([]entity.Photo, error) {
	photos, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return photos, err
}

func (s *PhotosImpl) Update(photo *entity.Photo, r *dto.AddPhotoRequest) (*entity.Photo, error) {
	photo, err := s.repository.Update(photo, r)
	if err != nil {
		return nil, err
	}

	return photo, err
}

func (s *PhotosImpl) Delete(photo *entity.Photo) error {
	err := s.repository.Delete(photo)
	if err != nil {
		return err
	}

	return nil
}
