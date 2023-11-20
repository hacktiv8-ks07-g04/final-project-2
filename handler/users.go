package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/dto"
	"github.com/hacktiv8-ks07-g04/final-project-2/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-2/service"
)

type Users interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type UsersImpl struct {
	userService service.Users
}

func NewUsers(userService service.Users) *UsersImpl {
	return &UsersImpl{
		userService: userService,
	}
}

func (h *UsersImpl) Register(c *gin.Context) {
	body := dto.RegisterRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		err := errs.New(http.StatusBadRequest, "Invalid request body!")
		c.Error(err)
		return
	}

	user, err := h.userService.Register(&body)
	if err != nil {
		err := errs.New(http.StatusInternalServerError, err.Error())
		c.Error(err)
		return
	}

	response := dto.UserResponse{
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

	token, err := h.userService.Login(&body)
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

	response := dto.LoginResponse{
		Token: token,
	}

	c.JSON(http.StatusOK, response)
}

func (h *UsersImpl) Update(c *gin.Context) {
	body := dto.UpdateUserRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		err := errs.New(http.StatusBadRequest, "Invalid request body!")
		c.Error(err)
		return
	}

	payload := dto.User{
		ID:       c.MustGet("userId").(uint),
		Username: body.Username,
		Email:    body.Email,
	}

	user, err := h.userService.Update(&payload)
	if err != nil {
		switch err.Error() {
		case "user not found":
			err := errs.New(http.StatusNotFound, err.Error())
			c.Error(err)
			return
		default:
			err := errs.New(http.StatusInternalServerError, err.Error())
			c.Error(err)
			return
		}
	}

	response := dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Age:       user.Age,
		UpdatedAt: &user.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func (h *UsersImpl) Delete(c *gin.Context) {
	header := c.MustGet("user").(map[string]interface{})
	userID := header["id"].(uint)

	err := h.userService.Delete(userID)
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
