package main

import (
	"fmt"
	"golang-auth/config"
	"golang-auth/internal/handlers"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.LoadDotenv()
	title := os.Getenv("SOME_TITLE")
	appEnv := os.Getenv("APP_ENV")
	fmt.Printf("This is title: %v, and env %v \n", title, appEnv)
	router.POST("/create-user", handlers.PostCreateUser)
	router.Run(":3000")
}
