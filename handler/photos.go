package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/errs"
	"github.com/hacktiv8-ks07-g04/final-project-2/service"
)

type Photos interface {
	Add(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type PhotosImpl struct {
	service service.Photos
}

func NewPhotos(service service.Photos) *PhotosImpl {
	return &PhotosImpl{service}
}

func (h *PhotosImpl) Add(c *gin.Context) {
	body := dto.AddPhotoRequest{}
	header := c.MustGet("user").(map[string]interface{})
	userID := header["id"].(uint)

	if err := c.ShouldBindJSON(&body); err != nil {
		err := errs.New(http.StatusBadRequest, "Invalid request body!")
		c.Error(err)
		return
	}

	photo, err := h.service.Add(userID, &body)
	if err != nil {
		err := errs.New(http.StatusInternalServerError, err.Error())
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
	photos, err := h.service.GetAll()
	if err != nil {
		err := errs.New(http.StatusInternalServerError, err.Error())
		c.Error(err)
		return
	}

	var response []dto.GetAllPhotosResponse

	for _, photo := range photos {
		response = append(response, dto.GetAllPhotosResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: dto.UserPhoto{
				Email:    photo.User.Email,
				Username: photo.User.Username,
			},
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *PhotosImpl) Update(c *gin.Context) {
	header := c.MustGet("user").(map[string]interface{})
	userID := header["id"].(uint)
	photoID := c.Param("photoId")

	if photoID == "" {
		err := errs.New(http.StatusBadRequest, "photo id is required in params")
		c.Error(err)
		return
	}

	body := dto.AddPhotoRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		err := errs.New(http.StatusBadRequest, "Invalid request body!")
		c.Error(err)
		return
	}

	photoIdInt, err := strconv.Atoi(photoID)
	if err != nil {
		err := errs.New(http.StatusBadRequest, "photo id must be a number")
		c.Error(err)
		return
	}

	photo, err := h.service.Update(uint(photoIdInt), userID, &body)
	if err != nil {
		if err.Error() == "photo not found" {
			err := errs.New(http.StatusNotFound, err.Error())
			c.Error(err)
			return
		} else if err.Error() == "you are not authorized to update this photo" {
			err := errs.New(http.StatusUnauthorized, err.Error())
			c.Error(err)
			return
		} else {
			err := errs.New(http.StatusInternalServerError, err.Error())
			c.Error(err)
			return
		}
	}

	response := dto.UpdatePhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func (h *PhotosImpl) Delete(c *gin.Context) {
	header := c.MustGet("user").(map[string]interface{})
	userID := header["id"].(uint)
	photoID := c.Param("photoId")

	if photoID == "" {
		err := errs.New(http.StatusBadRequest, "photo id is required in params")
		c.Error(err)
		return
	}

	photoIdInt, err := strconv.Atoi(photoID)
	if err != nil {
		err := errs.New(http.StatusBadRequest, "photo id must be a number")
		c.Error(err)
		return
	}

	if err := h.service.Delete(uint(photoIdInt), userID); err != nil {
		if err.Error() == "photo not found" {
			err := errs.New(http.StatusNotFound, err.Error())
			c.Error(err)
			return
		} else if err.Error() == "you are not authorized to update this photo" {
			err := errs.New(http.StatusUnauthorized, err.Error())
			c.Error(err)
			return
		} else {
			err := errs.New(http.StatusInternalServerError, err.Error())
			c.Error(err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
