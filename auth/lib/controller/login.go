package controller

import (
	// "errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"auth_service/lib/data"
	"auth_service/lib/jwt"

	"github.com/gin-gonic/gin"
)

// Note, this function should be private.
func getSignedToken() (string, error) {
	// Make JWT token with signing method of ES256 and claims.
	// Claims are attributes.
	// aud - audience
	// iss - issuer
	// exp - expiration of token
	claimsMap := map[string]string{
		"aud": "frontend.knowsearch.ml",
		"iss": "knowsearch.ml",
		"exp": fmt.Sprint(time.Now().Add(time.Minute * 1).Unix()),
	}

	// Provie shared secret. It should be very complex.
	// This should be passed as System Environment variable
	secret := os.Getenv("JWT_PRIVATE_KEY")
	header := "HS256"
	tokenString, err := jwt.GenerateJWT(header, claimsMap, secret)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

// Search user in database
func validateUser(email string, passHash string) (bool, error) {
	// usr, err := data.FindUserByEmail(email)
	// if err != nil {
	// 	return false, err
	// }

	// isPassValid := usr.ValidatePassHash(passHash)
	// if !isPassValid {
	// 	return false, nil
	// }

	return true, nil
}

func LoginHandler(ct *gin.Context) {
	// Validate request
	var userInput data.User
	if err := ct.ShouldBindJSON(&userInput); err != nil {
		ct.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	// if _, ok := req.Header["Email"]; !ok {
	// 	resW.WriteHeader(http.StatusBadRequest)
	// 	resW.Write([]byte("Email Missing."))
	// 	return
	// }

	// if _, ok := req.Header["PasswordHash"]; !ok {
	// 	resW.WriteHeader(http.StatusBadRequest)
	// 	resW.Write([]byte("PasswordHash Missing."))
	// 	return
	// }

	isPassValid, err := validateUser(userInput.Email, userInput.PasswordHash)
	if err != nil {
		// User not found
		ct.JSON(http.StatusUnauthorized, gin.H{"status": "User Not Found."})
		return
	}

	if !isPassValid {
		// Password was invalid
		ct.JSON(http.StatusUnauthorized, gin.H{"status": "Invalid Passord."})
		return
	}

	tokenString, err := getSignedToken()
	if err != nil {
		fmt.Println(err)
		ct.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error."})
		return
	}

	ct.JSON(http.StatusOK, gin.H{"jwt": tokenString})

}

// All requestes should be authenticated before accessing resource.
// func tokenValidater(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(resW http.ResponseWriter, req *http.Request) {
// 		// Check if token is present
// 		if _, ok := req.Header["Token"]; !ok {
// 			resW.WriteHeader(http.StatusUnauthorized)
// 			resW.Write([]byte("Token Missing"))
// 			return
// 		}

// 		token := req.Header["Token"][0]
// 		isValid, err := jwt.ValidateToken(token, os.Getenv("JWT_PRIVATE_KEY"))
	
// 		if err != nil {
// 			resW.WriteHeader(http.StatusInternalServerError)
// 			resW.Write([]byte("Token Validation Faild."))
// 			return
// 		}

// 		if !isValid {
// 			resW.WriteHeader(http.StatusUnauthorized)
// 			resW.Write([]byte("Token Invalid."))
// 			return
// 		}
// 		resW.WriteHeader(http.StatusOK)
// 		resW.Write([]byte("Authorized"))
// 	})
// }