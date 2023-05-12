// You shouldn’t export the whole user list outside this package
// to keep the module loosely coupled as microservice.

package data

import (
	"fmt"
	"os"

	// "log"
	"time"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email			string	`gorm:"size:255;not null;unique"`
	Username		string	`gorm:"size:255;not null;"`
	PasswordHash	string	`gorm:"size:255;not null"`
	Fullname		string

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
    // host := os.Getenv("ROACH_DB_HOST")
	// port := os.Getenv("ROACH_DB_PORT")
	// dbName := os.Getenv("ROACH_DB_DATABASE")
    // user := os.Getenv("ROACH_DB_USER")
    // password := os.Getenv("ROACH_DB_PASS")	

	dbName := os.Getenv("PGDATABASE")
	user := os.Getenv("PGUSER")
	pass := os.Getenv("PGPASSWORD")


	// roach_user := os.Getenv("ROACH_USER")
	// roach_pass := os.Getenv("ROACH_DB_PASS")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s",user, pass, dbName)
	// dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, dbName, user, password)

	// dsn := fmt.Sprintf("postgresql://%s:%s@sad-liger-4703.6xw.cockroachlabs.cloud:26257/authUser?sslmode=verify-full", roach_user, roach_pass)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to database!")
	}

	var now time.Time
	db.Raw("SELECT NOW()").Scan(&now)

	db.AutoMigrate(&User{})

	fmt.Println(now)
}



func FindUserByEmail(email string) (User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}


// Check if the password hash is valid
func (user *User) ValidatePassHash(pwdhash string) bool {
	return user.PasswordHash == pwdhash
}

// Add user
func AddUser(email string, username string, pass string, fullname string, role int) bool {
	// Query user data with condition
	var user User
	err := db.Where("Email = ? OR Username = ?", email, username).First(&user).Error
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





