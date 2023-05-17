// You shouldn’t export the whole user list outside this package
// to keep the module loosely coupled as microservice.

package data

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Email			string	`gorm:"size:255;not null;unique" json:"email"`
	Password		string	`gorm:"size:255;not null;" json:"-"` // json:"-" ensures password won't be returned in JSON response
	Role			int
}

type LoginUserInput struct {
	Email			string	`json:"email" binding:"required"`
	Password		string	`json:"password" binding:"required"`
}

// type LoginUserInput struct {
// 	User	NewUserInput	`gorm:"embedded"`
// }

type NewUserInput struct {
	Email			string	`json:"email" binding:"required"`
	Password		string	`json:"password" binding:"required"`
}

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

func Teardown() {
	migrator := db.Migrator()
	migrator.DropTable(&User{})
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

func FindUserById(id uint) (User, error) {
	var user User
	err := db.Where("ID=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}


// Add user
func (user *User) BeforeCreate (*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	// user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) CreateUser () (*User, error) {	
	err := db.Create(&user).Error
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

// Check if the password is valid
func (user *User) ValidatePass(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}



