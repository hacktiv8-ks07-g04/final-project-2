package server

import (
	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-2/handler"
	"github.com/hacktiv8-ks07-g04/final-project-2/infrastructure/database"
	"github.com/hacktiv8-ks07-g04/final-project-2/middleware"
	"github.com/hacktiv8-ks07-g04/final-project-2/repository"
	"github.com/hacktiv8-ks07-g04/final-project-2/service"
)

func Run() {
	server := Init()
	server.Run()
}

func Init() *gin.Engine {
	r := gin.Default()

	// Middleware
	r.Use(middleware.ErrorHandler())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Welcome to MyGram API",
		})
	})

	r.GET("/auth", middleware.Authentication(), func(c *gin.Context) {
		header := c.MustGet("user").(map[string]interface{})
		id := header["id"]

		c.JSON(200, gin.H{
			"status": "success",
			"header": header,
			"id":     id,
		})
	}) // for auth purpose

	db := database.GetInstance()

	// Users
	usersRepo := repository.NewUsers(db)
	usersService := service.NewUsers(usersRepo)
	usersHandler := handler.NewUsers(usersService)
	usersRouter := r.Group("/users")
	{
		usersRouter.POST("/register", usersHandler.Register)
		usersRouter.POST("/login", usersHandler.Login)
		usersRouter.PUT("/", middleware.Authentication(), usersHandler.Update)
		usersRouter.DELETE("/", middleware.Authentication(), usersHandler.Delete)
	}

	// Photos
	photosRepo := repository.NewPhotos(db)
	photosService := service.NewPhotos(photosRepo)
	photosHandler := handler.NewPhotos(photosService, usersService)
	photosRouter := r.Group("/photos").Use(middleware.Authentication())
	{
		photosRouter.POST("/", photosHandler.Add)
		photosRouter.GET("/", photosHandler.GetAll)
		photosRouter.PUT(
			"/:photoId",
			middleware.PhotoAuthorization(photosRepo),
			photosHandler.Update,
		)
		photosRouter.DELETE(
			"/:photoId",
			middleware.PhotoAuthorization(photosRepo),
			photosHandler.Delete,
		)
	}

	// Comments
	commentsRepo := repository.NewComments(db)
	commentsService := service.NewComments(commentsRepo)
	commentsHandler := handler.NewComments(commentsService)
	commentsRouter := r.Group("/comments").Use(middleware.Authentication())

	authService := service.NewAuthorization(commentsRepo)

	{
		commentsRouter.POST("/", commentsHandler.Add)
		commentsRouter.PUT(
			"/:commentId",
			authService.CommentAuthorization(),
			commentsHandler.Update,
		)
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
