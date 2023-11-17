package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/errs"
	"github.com/hacktiv8-ks07-g04/final-project-2/service"
)

type Users interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
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
		err := errs.New(http.StatusBadRequest, "Invalid request body!")
		c.Error(err)
		return
	}

	user, err := h.service.Register(&body)
	if err != nil {
		err := errs.New(http.StatusInternalServerError, err.Error())
		c.Error(err)
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

func (h *UsersImpl) Login(c *gin.Context) {
	body := dto.LoginRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		err := errs.New(http.StatusBadRequest, "Invalid request body!")
		c.Error(err)
		return
	}

	token, err := h.service.Login(&body)
	if err != nil {
		switch err.Error() {
		case "email or password is incorrect":
			err := errs.New(http.StatusUnauthorized, err.Error())
			c.Error(err)
			return
		case "failed to generate token":
			err := errs.New(http.StatusInternalServerError, err.Error())
			c.Error(err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *UsersImpl) Update(c *gin.Context) {
	body := dto.UpdateUserRequest{}
	header := c.MustGet("user").(map[string]interface{})
	userID := header["id"].(uint)

	if err := c.ShouldBindJSON(&body); err != nil {
		err := errs.New(http.StatusBadRequest, "Invalid request body!")
		c.Error(err)
		return
	}

	user, err := h.service.Update(userID, &body)
	if err != nil {
		switch err.Error() {
		case "user not found":
			err := errs.New(http.StatusNotFound, err.Error())
			c.Error(err)
			return
		case "failed to update user":
			err := errs.New(http.StatusInternalServerError, err.Error())
			c.Error(err)
			return
		}
	}

	response := dto.UpdateUserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func (h *UsersImpl) Delete(c *gin.Context) {
	header := c.MustGet("user").(map[string]interface{})
	userID := header["id"].(uint)

	err := h.service.Delete(userID)
	if err != nil {
		switch err.Error() {
		case "user not found":
			err := errs.New(http.StatusNotFound, err.Error())
			c.Error(err)
			return
		case "failed to update user":
			err := errs.New(http.StatusInternalServerError, err.Error())
			c.Error(err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
