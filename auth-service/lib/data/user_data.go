// You shouldn’t export the whole user list outside this package 
// to keep the module loosely coupled as microservice.

package data

import (
	"fmt"
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

// By default, GORM pluralizes struct name to snake_cases as table name
// snake_case as column name
// and uses CreatedAt, UpdatedAt to track creating/updating time
type User struct {
	// gorm.Model will include ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Email			string	`gorm:"size:255;not null;unique"`
	Username		string	`gorm:"size:255;not null;`
	// Password should be hashed on client 
	// In case of security breach, only hashed pass will be revealed
	PasswordHash	string	`gorm:"size:255;not null"`
	Fullname		string
	// Based on role level for authentication	
	// e.g. 0=standerd 1=admin
	Role			int
	// Age      		int64	`gorm:"column:age_of_the_beast"` // set name to `age_of_the_beast`
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
	// host		:= os.Getenv("DB_HOST")
	// username	:= os.Getenv("DB_USER")
	// password	:= os.Getenv("DB_PASSWORD")
	dbName		:= os.Getenv("DB_NAME")
	// port		:= os.Getenv("DB_PORT")

	// dsn := fmt.Sprintf("host=localhost password=%s dbname=%s", password, databaseName)
	dsn := fmt.Sprintf("dbname=%s", dbName)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to database!")
	}

	DB.AutoMigrate(&User{})
}



func FindUserByEmail(email string) (User, bool) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return User{}, false
	}

	return user, true
}


// Check if the password hash is valid
func (u *User) ValidatePassHash(pwdhash string) bool {
	return u.PasswordHash == pwdhash
}

// Add user
func AddUser(email string, username string, pass string, fullname string, role int) bool {
	// Query user data with condition
	var user User
	err := db.Where("Email = ?", email).First(&user).Error
	if err != nil {
		return false
	}
	
	newUser := User{
		Email:			email,
		Username:		username,
		PasswordHash:	pass,
		Fullname:		fullname,
		Role:			role,
	}

	db.Create(&newUser)
	return true
}





