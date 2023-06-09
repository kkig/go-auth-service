package controller

import (
	"auth/pkg/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Add a user to database
func RegisterUser(ctx *gin.Context) {
	// var userInput data.User
	var input data.NewUserInput

	// Extra error handling should be done on server to prevent malicious attacks
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	user := data.User{
		Email:			input.Email,
		Password:		input.Password,
		Role:			0,
	}

	savedUser, err := user.CreateUser()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// Validate, then add user
	// isAdded := data.AddUser(userInput.Email, userInput.Username, userInput.PasswordHash, userInput.Fullname, 0)


	ctx.JSON(http.StatusOK, gin.H{"user": savedUser})
}