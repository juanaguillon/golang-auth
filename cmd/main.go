package main

import (
	"golang-auth/config"
	"golang-auth/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.LoadDotenv()
	db, _ := config.DbLoad()
	defer db.Close()

	router.POST("/create-user", handlers.PostCreateUser)
	router.Run(":3000")
}
