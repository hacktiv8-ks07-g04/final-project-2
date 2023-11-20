package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-2/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-2/service"
)

type Comments interface {
	Add(c *gin.Context)
	Update(c *gin.Context)
}

type CommentsImpl struct {
	service service.Comments
}

func NewComments(service service.Comments) *CommentsImpl {
	return &CommentsImpl{service}
}

func (h *CommentsImpl) Add(c *gin.Context) {
	body := dto.AddCommentRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		err := errs.New(http.StatusBadRequest, "Invalid request body!")
		c.Error(err)
		return
	}

	payload := dto.Comment{
		Message: body.Message,
		PhotoID: body.PhotoID,
		UserID:  c.MustGet("userId").(uint),
	}

	comment, err := h.service.Add(&payload)
	if err != nil {
		err := errs.New(http.StatusInternalServerError, "Failed to add comment!")
		c.Error(err)
		return
	}

	response := dto.CommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: &comment.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func (h *CommentsImpl) Update(c *gin.Context) {
	body := dto.UpdateCommentRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		err := errs.New(http.StatusBadRequest, "Invalid request body!")
		c.Error(err)
		return
	}

	payload := c.MustGet("comment").(*entity.Comment)
	payload.Message = body.Message

	comment, err := h.service.Update(payload)
	if err != nil {
		err := errs.New(http.StatusInternalServerError, "Failed to update comment!")
		c.Error(err)
		return
	}

	response := dto.CommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		UpdatedAt: &comment.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}
