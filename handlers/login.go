package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Here, you would typically validate the username and password against your database or any other authentication mechanism
	// For demonstration purposes, let's assume a simple hardcoded check
	if loginData.Username == "user" && loginData.Password == "password" {
		// Authentication successful
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		// Authentication failed
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
