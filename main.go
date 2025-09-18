package main

import (
	"hoc-gin/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	r := gin.Default()

	userHandler := handlers.NewUserHandler()
	r.GET("/api/v1/users/:id", userHandler.GetUserByUuid)
	r.POST("/api/v1/users/", userHandler.CreateUser)

}
