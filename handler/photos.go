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
	GetAll(c *gin.Context)
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
	userId := c.MustGet("userId").(uint)

	var body dto.AddPhotoRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		err := errs.New(http.StatusBadRequest, "Invalid request body")
		c.Error(err)
		return
	}

	photo, err := h.photoService.Add(userId, &body)
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

	var response []dto.GetPhotoResponse
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
