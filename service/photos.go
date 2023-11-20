package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-2/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
)

type Photos interface {
	Add(payload *dto.Photo) (*entity.Photo, error)
	Get(id uint) (*entity.Photo, error)
	GetAll() ([]entity.Photo, error)
	Update(photo *entity.Photo, updatedData *dto.Photo) (*entity.Photo, error)
	Delete(photo *entity.Photo) error
}

type PhotosImpl struct {
	photoRepository repository.Photos
}

func NewPhotos(repository repository.Photos) *PhotosImpl {
	return &PhotosImpl{repository}
}

func (s *PhotosImpl) Add(payload *dto.Photo) (*entity.Photo, error) {
	photo := entity.Photo{
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoURL: payload.PhotoURL,
		UserID:   payload.UserID,
	}

	return s.photoRepository.Add(&photo)
}

func (s *PhotosImpl) Get(id uint) (*entity.Photo, error) {
	return s.photoRepository.Get(id)
}

func (s *PhotosImpl) GetAll() ([]entity.Photo, error) {
	return s.photoRepository.GetAll()
}

func (s *PhotosImpl) Update(photo *entity.Photo, updatedData *dto.Photo) (*entity.Photo, error) {
	return s.photoRepository.Update(photo, updatedData)
}

func (s *PhotosImpl) Delete(photo *entity.Photo) error {
	return s.photoRepository.Delete(photo)
}
