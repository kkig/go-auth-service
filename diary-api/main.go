package main

import (
	"github.com/db"
	"github.com/model"

	"log"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDB()
}

func loadDB() {
	db.Connect()
	db.DB.AutoMigrate(&model.User{})
	db.DB.AutoMigrate(&model.Entry{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
}