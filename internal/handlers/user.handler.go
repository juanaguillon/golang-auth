package handlers

import (
	"golang-auth/internal/models"
	"golang-auth/pkg"
	"net/http"

	// "golang-auth/internal/repository"
	// "golang-auth/internal/service"

	"github.com/gin-gonic/gin"
)

func PostCreateUser(c *gin.Context) {
	var createUserBody models.UserCreateRequest

	if err := pkg.DisallowUnkownJSONProps(c.Request.Body, &createUserBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := c.ShouldBindBodyWithJSON(&createUserBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createUserBody})
}
