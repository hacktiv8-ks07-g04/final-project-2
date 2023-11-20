package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-2/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
)

type Photos interface {
	Add(userID uint, r *dto.AddPhotoRequest) (*entity.Photo, error)
	GetAll() ([]entity.Photo, error)
	Update(photoId uint, r *dto.UpdatePhotoRequest) (*entity.Photo, error)
	Delete(photoId uint) error
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
		UserID:   userID,
	}

	result, err := s.repository.Add(&photo)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *PhotosImpl) GetAll() ([]entity.Photo, error) {
	result, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *PhotosImpl) Update(photoId uint, r *dto.UpdatePhotoRequest) (*entity.Photo, error) {
	result, err := s.repository.Update(photoId, r)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *PhotosImpl) Delete(photoId uint) error {
	err := s.repository.Delete(photoId)
	if err != nil {
		return err
	}

	return nil
}
