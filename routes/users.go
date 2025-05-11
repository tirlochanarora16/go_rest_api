package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

type response gin.H

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse the requested data",
			"error":   err.Error(),
		})
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not save the user",
			"error":   err.Error(),
		})
	}

	context.JSON(http.StatusCreated, response{
		"message": "user created successfully",
	})
}
