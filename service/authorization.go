package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
)

type Authorization interface {
	CommentAuthorization(r *repository.CommentsImpl) gin.HandlerFunc
}

type AuthorizationImpl struct {
	commentRepository repository.Comments
}

func NewAuthorization(commentRepository repository.Comments) *AuthorizationImpl {
	return &AuthorizationImpl{
		commentRepository: commentRepository,
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
