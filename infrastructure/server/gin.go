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
	})

	db := database.GetInstance()

	// Users
	usersRepo := repository.NewUsers(db)
	usersService := service.NewUsers(usersRepo)
	usersHandler := handler.NewUsers(usersService)

	// Photos
	photosRepo := repository.NewPhotos(db)
	photosService := service.NewPhotos(photosRepo)
	photosHandler := handler.NewPhotos(photosService, usersService)

	// Comments
	commentsRepo := repository.NewComments(db)
	commentsService := service.NewComments(commentsRepo)
	commentsHandler := handler.NewComments(commentsService, photosService, usersService)

	// Social Medias
	socialMediasRepo := repository.NewSocialMedias(db)
	socialMediasService := service.NewSocialMedias(socialMediasRepo)
	socialMediasHandler := handler.NewSocialMedias(socialMediasService)

	// Authorization
	authService := service.NewAuthorization(photosRepo, commentsRepo)

	// Routes
	usersRouter := r.Group("/users")
	{
		usersRouter.POST("/register", usersHandler.Register)
		usersRouter.POST("/login", usersHandler.Login)
		usersRouter.PUT("/", middleware.Authentication(), usersHandler.Update)
		usersRouter.DELETE("/", middleware.Authentication(), usersHandler.Delete)
	}

	photosRouter := r.Group("/photos").Use(middleware.Authentication())
	{
		photosRouter.POST("/", photosHandler.Add)
		photosRouter.GET("/", photosHandler.GetAll)
		photosRouter.PUT(
			"/:photoId",
			authService.PhotoAuthorization(),
			photosHandler.Update,
		)
		photosRouter.DELETE(
			"/:photoId",
			authService.PhotoAuthorization(),
			photosHandler.Delete,
		)
	}

	commentsRouter := r.Group("/comments").Use(middleware.Authentication())
	{
		commentsRouter.POST("/", commentsHandler.Add)
		commentsRouter.GET("/", commentsHandler.GetAll)
		commentsRouter.PUT(
			"/:commentId",
			authService.CommentAuthorization(),
			commentsHandler.Update,
		)
		commentsRouter.DELETE(
			"/:commentId",
			authService.CommentAuthorization(),
			commentsHandler.Delete,
		)
	}

	socialMediasRouter := r.Group("/socialmedias").Use(middleware.Authentication())
	{
		socialMediasRouter.POST("/", socialMediasHandler.Add)
	}

	return r
}
