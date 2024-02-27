package middleware

import (
	"github.com/gin-gonic/gin"
)

func Midware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Middleware logic
	}
}
