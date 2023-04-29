package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	// host := os.Getenv("DB_HOST")
	// username := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	// port := os.Getenv("DB_PORT")	

	// dsn := fmt.Sprintf("host=localhost password=%s dbname=%s", password, databaseName)
	dsn := fmt.Sprintf("dbname=%s",databaseName)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to database!")
	}
}