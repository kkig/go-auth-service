package main

import (
	"auth_service/lib/data"
	// "auth_service/lib/controller"

	"fmt"
	"log"
	// "os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	data.Connect()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load("./.env.local")
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
}

func serveApplication() {
	port := ":8000"
	router := gin.Default()

	v1 := router.Group("/auth")
	{
		// v1.POST("/register", controller.RegisterHandler)
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

