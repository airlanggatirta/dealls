package routes

import (
	"dealls/handlers"
	"dealls/middleware"
	"dealls/repository"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userRepository *repository.UserRepository, premiumRepository *repository.PremiumRepository) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Guys!",
		})
	})

	// Routes for authentication
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", handlers.LoginHandler(userRepository))
		authGroup.POST("/register", middleware.Midware(), handlers.RegisterHandler(userRepository, premiumRepository))
		authGroup.POST("/check-status", handlers.CheckPremiumUserHandler(userRepository))
	}

	premiumGroup := router.Group("/premium")
	{
		premiumGroup.GET("/get-list", middleware.Midware(), handlers.GetPremiumList(premiumRepository))
	}

	// router.POST("/logout", middleware.Midware(), handlers.LogoutHandler)

	// router.POST("/make-payment", middleware.Midware(), handlers.LoginHandler)

	// router.POST("/get-swipe-list", middleware.Midware(), handlers.LoginHandler)
	// router.POST("/pass", middleware.Midware(), handlers.LoginHandler)
	// router.POST("/match", middleware.Midware(), handlers.LoginHandler)

}
