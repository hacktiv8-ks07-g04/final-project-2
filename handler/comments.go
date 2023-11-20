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
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type CommentsImpl struct {
	commentService service.Comments
	userService    service.Users
	photoService   service.Photos
}

func NewComments(
	commentService service.Comments,
	photoService service.Photos,
	userService service.Users,
) *CommentsImpl {
	return &CommentsImpl{
		commentService: commentService,
		photoService:   photoService,
		userService:    userService,
	}
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

	comment, err := h.commentService.Add(&payload)
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

func (h *CommentsImpl) GetAll(c *gin.Context) {
	comments, err := h.commentService.GetAll()
	if err != nil {
		err := errs.New(http.StatusInternalServerError, "Failed to get comments!")
		c.Error(err)
		return
	}

	response := []dto.CommentResponse{}

	for _, comment := range comments {
		var err error
		user, err := h.userService.Get(comment.UserID)
		if err != nil {
			err := errs.New(http.StatusInternalServerError, "Failed to get user!")
			c.Error(err)
			return
		}

		photo, err := h.photoService.Get(comment.PhotoID)
		if err != nil {
			err := errs.New(http.StatusInternalServerError, "Failed to get photo!")
			c.Error(err)
			return
		}

		response = append(response, dto.CommentResponse{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			CreatedAt: &comment.CreatedAt,
			UpdatedAt: &comment.UpdatedAt,
			User: &dto.UserResponse{
				ID:       user.ID,
				Username: user.Username,
				Email:    user.Email,
			},
			Photo: &dto.PhotoResponse{
				ID:       photo.ID,
				Title:    photo.Title,
				Caption:  photo.Caption,
				PhotoURL: photo.PhotoURL,
				UserID:   photo.UserID,
			},
		})
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

	comment := c.MustGet("comment").(*entity.Comment)

	payload := &dto.Comment{
		Message: body.Message,
	}

	newComment, err := h.commentService.Update(comment, payload)
	if err != nil {
		err := errs.New(http.StatusInternalServerError, "Failed to update comment!")
		c.Error(err)
		return
	}

	response := dto.CommentResponse{
		ID:        newComment.ID,
		Message:   newComment.Message,
		PhotoID:   newComment.PhotoID,
		UserID:    newComment.UserID,
		UpdatedAt: &newComment.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func (h *CommentsImpl) Delete(c *gin.Context) {
	payload := c.MustGet("comment").(*entity.Comment)
	err := h.commentService.Delete(payload)
	if err != nil {
		err := errs.New(http.StatusInternalServerError, "Failed to delete comment!")
		c.Error(err)
		return
	}

	response := dto.DeleteCommentResponse{
		Message: "Your comment has been successfully deleted",
	}

	c.JSON(http.StatusOK, response)
}
