package middlwares

import (
	"net/http"

	"example.com/exercises/events-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	authToken := c.Request.Header.Get("Authorization")
	if authToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})

		return
	}

	userId, err := utils.VerifyToken(authToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})

		return
	}

	c.Set("UserId", userId)
	c.Next()
}
