package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-2/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-2/service"
)

type Photos interface {
	Add(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type PhotosImpl struct {
	photoService service.Photos
	userService  service.Users
}

func NewPhotos(photoService service.Photos, userService service.Users) *PhotosImpl {
	return &PhotosImpl{
		photoService: photoService,
		userService:  userService,
	}
}

func (h *PhotosImpl) Add(c *gin.Context) {
	body := dto.AddPhotoRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		err := errs.New(http.StatusBadRequest, "Invalid request body")
		c.Error(err)
		return
	}

	payload := dto.Photo{
		Title:    body.Title,
		Caption:  body.Caption,
		PhotoURL: body.PhotoURL,
		UserID:   c.MustGet("userId").(uint),
	}

	photo, err := h.photoService.Add(&payload)
	if err != nil {
		c.Error(err)
		return
	}

	response := dto.AddPhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *PhotosImpl) GetAll(c *gin.Context) {
	photos, err := h.photoService.GetAll()
	if err != nil {
		err := errs.New(http.StatusNotFound, "Photos not found")
		c.Error(err)
		return
	}

	response := []dto.GetPhotoResponse{}

	for _, photo := range photos {
		user, err := h.userService.Get(photo.UserID)
		if err != nil {
			err := errs.New(http.StatusNotFound, "User not found")
			c.Error(err)
			return
		}

		response = append(response, dto.GetPhotoResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
			User: dto.UserPhoto{
				Email:    user.Email,
				Username: user.Username,
			},
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *PhotosImpl) Update(c *gin.Context) {
	body := dto.UpdatePhotoRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		err := errs.New(http.StatusBadRequest, "Invalid request body!")
		c.Error(err)
		return
	}

	photo := c.MustGet("photo").(*entity.Photo)

	payload := dto.Photo{
		Title:    body.Title,
		Caption:  body.Caption,
		PhotoURL: body.PhotoURL,
	}

	newPhoto, err := h.photoService.Update(photo, &payload)
	if err != nil {
		if err.Error() == "photo not found" {
			err := errs.New(http.StatusNotFound, err.Error())
			c.Error(err)
			return
		} else {
			err := errs.New(http.StatusInternalServerError, err.Error())
			c.Error(err)
			return
		}
	}

	response := dto.PhotoResponse{
		ID:        newPhoto.ID,
		Title:     newPhoto.Title,
		Caption:   newPhoto.Caption,
		PhotoURL:  newPhoto.PhotoURL,
		UserID:    newPhoto.UserID,
		UpdatedAt: &newPhoto.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func (h *PhotosImpl) Delete(c *gin.Context) {
	photo := c.MustGet("photo").(*entity.Photo)

	if err := h.photoService.Delete(photo); err != nil {
		err := errs.New(http.StatusInternalServerError, err.Error())
		c.Error(err)
		return
	}

	response := dto.DeletePhotoResponse{
		Message: "Your photo has been successfully deleted",
	}

	c.JSON(http.StatusOK, response)
}
