package server

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	server := Init()
	server.Run()
}

func Init() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Welcome to MyGram API",
		})
	})

	return r
}
