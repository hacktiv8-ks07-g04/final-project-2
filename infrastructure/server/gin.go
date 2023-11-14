package server

import (
	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/handler"
	"github.com/hacktiv8-ks07-g04/final-project-2/infrastructure/database"
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
	"github.com/hacktiv8-ks07-g04/final-project-2/service"
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

	db := database.GetInstance()

	// Users
	usersRepo := repository.NewUsers(db)
	usersService := service.NewUsers(usersRepo)
	usersHandler := handler.NewUsers(usersService)
	usersRouter := r.Group("/users")
	{
		usersRouter.POST("/register", usersHandler.Register)
	}

	// Photos
	photosRepo := repository.NewPhotos(db)
	photosService := service.NewPhotos(photosRepo)
	photosHandler := handler.NewPhotos(photosService)
	photosRouter := r.Group("/photos")
	{
		_, _ = photosHandler, photosRouter
	}

	// Comments
	commentsRepo := repository.NewComments(db)
	commentsService := service.NewComments(commentsRepo)
	commentsHandler := handler.NewComments(commentsService)
	commentsRouter := r.Group("/comments")
	{
		_, _ = commentsHandler, commentsRouter
	}

	// Social Medias
	socialMediasRepo := repository.NewSocialMedias(db)
	socialMediasService := service.NewSocialMedias(socialMediasRepo)
	socialMediasHandler := handler.NewSocialMedias(socialMediasService)
	socialMediasRouter := r.Group("/socialmedias")
	{
		_, _ = socialMediasHandler, socialMediasRouter
	}

	return r
}
