package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/errs"
	"github.com/hacktiv8-ks07-g04/final-project-2/service"
)

type Photos interface {
	Add(c *gin.Context)
}

type PhotosImpl struct {
	service service.Photos
}

func NewPhotos(service service.Photos) *PhotosImpl {
	return &PhotosImpl{service}
}

func (h *PhotosImpl) Add(c *gin.Context) {
	userId := c.MustGet("userId").(uint)

	var body dto.AddPhotoRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		err := errs.New(http.StatusBadRequest, "Invalid request body")
		c.Error(err)
		return
	}

	photo, err := h.service.Add(userId, &body)
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
