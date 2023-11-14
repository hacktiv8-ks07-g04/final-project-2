package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/service"
)

type Users interface {
	Register(c *gin.Context)
}

type UsersImpl struct {
	service service.Users
}

func NewUsers(service service.Users) *UsersImpl {
	return &UsersImpl{service}
}

func (h *UsersImpl) Register(c *gin.Context) {
	body := dto.RegisterRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "invalid request body",
		})
		return
	}

	user, err := h.service.Register(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	response := dto.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}

	c.JSON(http.StatusCreated, response)
}
