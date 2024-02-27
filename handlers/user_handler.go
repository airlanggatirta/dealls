package handlers

import (
	"net/http"

	"dealls/repository"
	"dealls/services"

	"github.com/gin-gonic/gin"
)

func LoginHandler(userRepository *repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginData struct {
			UserEmail string `json:"email" binding:"required"`
			Password  string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userService := services.NewUserService(userRepository, nil)

		if err := userService.Login(loginData.UserEmail, loginData.Password); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	}
}

func RegisterHandler(userRepository *repository.UserRepository, premiumRepository *repository.PremiumRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var registrationData struct {
			Username           string `json:"username" binding:"required"`
			UserEmail          string `json:"email" binding:"required,email"`
			UserPassword       string `json:"password" binding:"required"`
			UserFullname       string `json:"fullname" binding:"required"`
			UserProfilePicture string `json:"profile_picture"`
		}

		if err := c.ShouldBindJSON(&registrationData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userService := services.NewUserService(userRepository, premiumRepository)

		if err := userService.Register(registrationData.Username, registrationData.UserPassword, registrationData.UserEmail, registrationData.UserFullname, registrationData.UserProfilePicture); err != nil {
			if err == services.ErrUsernameExists {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	}
}

func CheckPremiumUserHandler(userRepository *repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userData struct {
			UserID uint64 `json:"user_id" binding:"required"`
		}

		if err := c.ShouldBindJSON(&userData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userService := services.NewUserService(userRepository, nil)

		user, err := userService.CheckPremiumUserStatus(userData.UserID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
