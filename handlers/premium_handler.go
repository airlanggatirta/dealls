package handlers

import (
	"fmt"
	"net/http"

	"dealls/repository"
	"dealls/services"

	"github.com/gin-gonic/gin"
)

func GetPremiumList(premiumRepository *repository.PremiumRepository) gin.HandlerFunc {
	return func(c *gin.Context) {

		premiumService := services.NewPremiumService(premiumRepository, nil)

		paramPairs := c.Request.URL.Query()
		for key, values := range paramPairs {
			fmt.Printf("key = %v, value(s) = %v\n", key, values)
		}

		if err := premiumService.GetPremium(1); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}
	}
}
