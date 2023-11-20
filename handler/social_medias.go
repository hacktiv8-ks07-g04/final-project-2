package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-2/service"
)

type SocialMedias interface {
	Add(c *gin.Context)
}

type SocialMediasImpl struct {
	service service.SocialMedias
}

func NewSocialMedias(service service.SocialMedias) *SocialMediasImpl {
	return &SocialMediasImpl{service}
}

func (h *SocialMediasImpl) Add(c *gin.Context) {
	body := dto.AddSocialMediaRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		err := errs.New(http.StatusBadRequest, "Invalid request body!")
		c.Error(err)
		return
	}

	payload := dto.SocialMedia{
		Name:           body.Name,
		SocialMediaURL: body.SocialMediaURL,
		UserID:         c.MustGet("userId").(uint),
	}

	socialMedia, err := h.service.Add(&payload)
	if err != nil {
		err := errs.New(http.StatusInternalServerError, "Failed to add social media")
		c.Error(err)
		return
	}

	response := dto.SocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
		CreatedAt:      &socialMedia.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}
