package handlers

import (
	"golang-auth/internal/models"
	"golang-auth/pkg"
	"net/http"

	"golang-auth/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandlerIface interface {
	PostCreateUser(c *gin.Context)
}

type UserHandler struct {
	userService service.UserServiceIface
}

func (uhandler *UserHandler) PostCreateUser(c *gin.Context) {
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

	uhandler.userService.CreateUser(createUserBody)

	c.JSON(http.StatusCreated, gin.H{"data": createUserBody})
}
