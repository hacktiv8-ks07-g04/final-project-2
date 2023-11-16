package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/errs"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err != nil {
			switch e := err.Err.(type) {
			case *errs.CustomError:
				c.AbortWithStatusJSON(e.Code, gin.H{
					"error": e.Message,
				})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}
		}
	}
}
