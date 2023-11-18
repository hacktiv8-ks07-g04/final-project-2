package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/errs"
	"github.com/hacktiv8-ks07-g04/final-project-2/service"
)

type Comments interface {
	Add(c *gin.Context)
	GetAll(c *gin.Context)
}

type CommentsImpl struct {
	service service.Comments
}

func NewComments(service service.Comments) *CommentsImpl {
	return &CommentsImpl{service}
}

func (h *CommentsImpl) Add(c *gin.Context) {
	header := c.MustGet("user").(map[string]interface{})
	userID := header["id"].(uint)

	body := dto.CreateCommentRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		err := errs.New(http.StatusBadRequest, "Invalid request body!")
		c.Error(err)
		return
	}

	comment, err := h.service.Add(userID, &body)
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

	response := dto.CreateCommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		UserID:    comment.UserID,
		PhotoID:   comment.PhotoID,
		CreatedAt: comment.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *CommentsImpl) GetAll(c *gin.Context) {
	comments, err := h.service.GetAll()
	if err != nil {
		err := errs.New(http.StatusInternalServerError, err.Error())
		c.Error(err)
		return
	}

	var response []dto.GetAllCommentsResponse

	for _, comment := range comments {
		response = append(response, dto.GetAllCommentsResponse{
			ID:        comment.ID,
			Message:   comment.Message,
			UserID:    comment.UserID,
			PhotoID:   comment.PhotoID,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			User: dto.UserComment{
				ID:       comment.User.ID,
				Email:    comment.User.Email,
				Username: comment.User.Username,
			},
			Photo: dto.PhotoComment{
				ID:       comment.Photo.ID,
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoURL: comment.Photo.PhotoURL,
				UserID:   comment.Photo.UserID,
			},
		})
	}

	c.JSON(http.StatusOK, response)
}
