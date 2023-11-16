package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/domain/errs"
	"github.com/hacktiv8-ks07-g04/final-project-2/utils"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := utils.VerifyToken(c)
		if err != nil {
			err := errs.New(http.StatusUnauthorized, err.Error())
			c.Error(err)
		}

		c.Set("user", claims)
		c.Next()
	}
}
