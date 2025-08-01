package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadDotenv() {
	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "development"
	}

	godotenv.Load(".env." + env + ".local")
	if "test" != env {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load() // The Original .env
}
