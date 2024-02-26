package middleware

import (
	"github.com/gin-gonic/gin"
)

func Midware() gin.HandlerFunc {
	// Your middleware logic for middleware 1
	return func(c *gin.Context) {
		// Middleware logic
	}
}
