// You shouldn’t export the whole user list outside this package
// to keep the module loosely coupled as microservice.

package data

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Email			string	`gorm:"size:255;not null;unique" json:"email"`
	PasswordHash	string	`gorm:"size:255;" json:"-"` // json:"-" ensures password won't be returned in JSON response
	Role			int
}

type NewUserInput struct {
	Email			string	`json:"email" binding:"required"`
	PasswordHash	string	`json:"password" binding:"required"`
}

// var userList = []user{
// 	{
// 		email:        "abc@gmail.com",
// 		username:     "abc12",
// 		passwordhash: "hashedme1",
// 		fullname:     "abc def",
// 		createDate:   "1631600786",
// 		role:         1,
// 	},
// 	{
// 		email:        "chekme@example.com",
// 		username:     "checkme34",
// 		passwordhash: "hashedme2",
// 		fullname:     "check me",
// 		createDate:   "1631600837",
// 		role:         0,
// 	},
// }

var db *gorm.DB

func Connect() {
    var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
		  SlowThreshold:              time.Second,   // Slow SQL threshold
		  LogLevel:                   logger.Silent, // Log level
		  IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
		  ParameterizedQueries:      true,           // Don't include params in the SQL log
		  Colorful:                  false,          // Disable color
		},
	)
	
	dbName := os.Getenv("PGDATABASE")
    host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	pass := os.Getenv("PGPASSWORD")

	// dsn := fmt.Sprintf("user=%s dbname=%s",user, dbName)
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, dbName, user, pass)

	// dsn := fmt.Sprintf("postgresql://%s:%s@sad-liger-4703.6xw.cockroachlabs.cloud:26257/authUser?sslmode=verify-full", roach_user, roach_pass)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger,})
	
	if err != nil {
		panic(err)
	}

	// Migrate schema change if any - Create table if it doesn't exist.
	db.AutoMigrate(&User{})

	fmt.Println("Connected to server!")

}

func FindUserByEmail(email string) (User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	// err := db.Where("Email = ? OR Username = ?", email, username).First(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}


// Add user
func (user *User) AddUser () (*User, error) {	
	err := db.Create(&user).Error

	if err != nil {
		return &User{}, err
	}

	return user, nil
}

// Check if the password hash is valid
func (user *User) ValidatePassHash(pwdhash string) bool {
	return user.PasswordHash == pwdhash
}



