package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/errs"
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
)

// user's own photo check middleware
func PhotoAuthorization(r *repository.PhotosImpl) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.MustGet("user").(map[string]interface{})
		userID := header["id"].(uint)
		photoID, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			err := errs.New(http.StatusBadRequest, "photo id must be a number")
			c.Error(err)
			return
		}

		photo, err := r.Get(uint(photoID))
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

		// pass the photo to the next middleware
		c.Set("photo", photo)

		c.Next()
	}
}

func CommentAuthorization(r *repository.CommentsImpl) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.MustGet("user").(map[string]interface{})
		userID := header["id"].(uint)
		commentID, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			err := errs.New(http.StatusBadRequest, "comment id must be a number")
			c.Error(err)
			c.Abort()
		}

		comment, err := r.Get(uint(commentID))
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
			err := errs.New(http.StatusUnauthorized, "you are not authorized to update this photo")
			c.Error(err)
			c.Abort()
			return
		}

		c.Set("comment", comment)

		c.Next()
	}
}
