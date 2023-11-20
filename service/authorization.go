package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
)

type Authorization interface {
	PhotoAuthorization(r *repository.PhotosImpl) gin.HandlerFunc
	CommentAuthorization(r *repository.CommentsImpl) gin.HandlerFunc
	SocialMediaAuthorization(r *repository.SocialMediasImpl) gin.HandlerFunc
}

type AuthorizationImpl struct {
	photoRepository       repository.Photos
	commentRepository     repository.Comments
	socialMediaRepository repository.SocialMedias
}

func NewAuthorization(
	photoRepository repository.Photos,
	commentRepository repository.Comments,
	socialMediaRepository repository.SocialMedias,
) *AuthorizationImpl {
	return &AuthorizationImpl{
		photoRepository:       photoRepository,
		commentRepository:     commentRepository,
		socialMediaRepository: socialMediaRepository,
	}
}

func (s *AuthorizationImpl) PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("userId").(uint)
		photoID, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			err := errs.New(http.StatusBadRequest, "invalid photo id")
			c.Error(err)
			return
		}

		photo, err := s.photoRepository.Get(uint(photoID))
		if err != nil {
			if err.Error() == "photo not found" {
				err := errs.New(http.StatusNotFound, err.Error())
				c.Error(err)
				c.Abort()
				return
			} else {
				c.Error(err)
				c.Abort()
				return
			}
		}

		if photo.UserID != userID {
			err := errs.New(http.StatusUnauthorized, "you are not authorized to update this photo")
			c.Error(err)
			c.Abort()
			return
		}

		c.Set("photo", photo)

		c.Next()
	}
}

func (s *AuthorizationImpl) CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("userId").(uint)
		commentID, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			err := errs.New(http.StatusBadRequest, "invalid comment id")
			c.Error(err)
			return
		}

		comment, err := s.commentRepository.Get(uint(commentID))
		if err != nil {
			if err.Error() == "comment not found" {
				err := errs.New(http.StatusNotFound, err.Error())
				c.Error(err)
				c.Abort()
				return
			} else {
				c.Error(err)
				c.Abort()
				return
			}
		}

		if comment.UserID != userID {
			err := errs.New(
				http.StatusUnauthorized,
				"you are not authorized to update this comment",
			)
			c.Error(err)
			c.Abort()
			return
		}

		c.Set("comment", comment)
		c.Next()
	}
}

func (s *AuthorizationImpl) SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("userId").(uint)
		socialMedaiId, err := strconv.Atoi(c.Param("socialMediaId"))
		if err != nil {
			err := errs.New(http.StatusBadRequest, "invalid social media id")
			c.Error(err)
			return
		}

		socialMedia, err := s.socialMediaRepository.Get(uint(socialMedaiId))
		if err != nil {
			if err.Error() == "social media not found" {
				err := errs.New(http.StatusNotFound, err.Error())
				c.Error(err)
				c.Abort()
				return
			} else {
				c.Error(err)
				c.Abort()
				return
			}
		}

		if socialMedia.UserID != userID {
			err := errs.New(
				http.StatusUnauthorized,
				"you are not authorized to update this comment",
			)
			c.Error(err)
			c.Abort()
			return
		}

		c.Set("socialMedia", socialMedia)
		c.Next()
	}
}
