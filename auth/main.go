package main

import (
	"auth_service/lib/data"
	"auth_service/lib/controller"

	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"log"
)

func main() {
	loadEnv()
	data.Connect()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	port := ":8080"
	router := gin.Default()

	v1 := router.Group("/auth")
	{
		v1.POST("/register", controller.RegisterUser)
		// v1.POST("/login", controller.LoginHandler)
	
		// Test
		v1.GET("/ping", func(cx *gin.Context) {
			cx.JSON(200, gin.H{
				"message": "pong!",
			})
		})		
	}

	router.Run(port)
	fmt.Println("Server running on " + port)
}

