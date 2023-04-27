package main

import (
	"belajar-oauth/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/google/login", handlers.GoogleLogin)
	r.GET("/google/callback", handlers.GoogleCallback)

	r.Run()
}
