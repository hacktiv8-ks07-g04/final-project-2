package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/utils"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := utils.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}

		user := claims.(map[string]interface{})
		c.Set("user", user)
		c.Set("userId", user["id"])

		c.Next()
	}
}
