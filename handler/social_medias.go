package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-2/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-2/service"
)

type SocialMedias interface {
	Add(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
}

type SocialMediasImpl struct {
	socialMediaService service.SocialMedias
	userService        service.Users
}

func NewSocialMedias(socialMediaService service.SocialMedias, userService service.Users) *SocialMediasImpl {
	return &SocialMediasImpl{
		socialMediaService: socialMediaService,
		userService:        userService,
	}
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

	socialMedia, err := h.socialMediaService.Add(&payload)
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

func (h *SocialMediasImpl) GetAll(c *gin.Context) {
	socialMedias, err := h.socialMediaService.GetAll()
	if err != nil {
		err := errs.New(http.StatusInternalServerError, "Invalid request body!")
		c.Error(err)
		return
	}

	response := []dto.SocialMediaResponse{}

	for _, socialMedia := range socialMedias {
		user, err := h.userService.Get(socialMedia.UserID)
		if err != nil {
			err := errs.New(http.StatusInternalServerError, "Failed to get user!")
			c.Error(err)
			return
		}

		response = append(response, dto.SocialMediaResponse{
			ID:             socialMedia.ID,
			Name:           socialMedia.Name,
			SocialMediaURL: socialMedia.SocialMediaURL,
			UserID:         socialMedia.UserID,
			CreatedAt:      &socialMedia.CreatedAt,
			UpdatedAt:      &socialMedia.UpdatedAt,
			User: &dto.UserResponse{
				ID:       user.ID,
				Username: user.Username,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"social_medias": response,
	})
}

func (h *SocialMediasImpl) Update(c *gin.Context) {
	body := dto.UpdateSocialMediaRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		err := errs.New(http.StatusBadRequest, "Invalid request body!")
		c.Error(err)
		return
	}

	payload := c.MustGet("socialMedia").(*entity.SocialMedia)
	payload.Name = body.Name
	payload.SocialMediaURL = body.SocialMediaURL

	socialMedia, err := h.socialMediaService.Update(payload)
	if err != nil {
		err := errs.New(http.StatusInternalServerError, "Failed to update comment!")
		c.Error(err)
		return
	}

	response := dto.SocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
		UpdatedAt:      &socialMedia.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}
