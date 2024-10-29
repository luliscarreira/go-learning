package routes

import (
	"fmt"
	"net/http"

	"example.com/exercises/events-api/models"
	"example.com/exercises/events-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not process request",
		})

		return
	}

	err = user.Save()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not process request",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func login(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not process request",
		})

		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not authenticate user",
		})

		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not authenticate user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User authenticated successfully",
		"token":   token,
	})
}
