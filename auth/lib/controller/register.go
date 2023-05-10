package controller

import (
	"net/http"

	"auth_service/lib/data"
	"github.com/gin-gonic/gin"
)

// Add a user to database
func RegisterUser(ct *gin.Context) {
	var userInput data.User

	// Extra error handling should be done on server to prevent malicious attacks
	if err := ct.ShouldBindJSON(&userInput); err != nil {
		ct.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	// Validate, then add user
	isAdded := data.AddUser(userInput.Email, userInput.Username, userInput.PasswordHash, userInput.Fullname, 0)

	// User was already in database
	if !isAdded {
		ct.JSON(http.StatusConflict, gin.H{"status": "Email or Username aldready exists."})
		return
	}

	ct.JSON(http.StatusOK, gin.H{"status": "User Created."})
}