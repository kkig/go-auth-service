package controller

import (
	"net/http"

	"auth/pkg/data"
	"auth/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func LoginUser(ctx *gin.Context) {
	// Validate request
	var input data.LoginUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	user, err := data.FindUserByEmail(input.Email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePass(input.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := jwt.GenerateJWT(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	ctx.JSON(http.StatusOK, gin.H{"jwt": jwt})

}


// Test for protected API - require pre-authentication
func TestProtected(ctx *gin.Context) {
	_, err := jwt.LookupUserByUID(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// result := map[string]interface{} {
	// 	"id":		user.ID,
	// 	"email": 	user.Email,
	// }

	ctx.JSON(http.StatusOK, gin.H{"message": "authenticated user"})
}