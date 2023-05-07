package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"
	"os"

	"auth_service/lib/data"
	"auth_service/lib/jwt"
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
	tokenString, err := jwt.GenerateToken(header, claimsMap, secret)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

// Search user in database
func validateUser(email string, passHash string) (bool, error) {
	usr, isFound := data.FindUserByEmail(email)
	if !isFound {
		return false, errors.New("User Not Found.")
	}

	isPassValid := usr.ValidatePassHash(passHash)
	if !isPassValid {
		return false, nil
	}

	return true, nil
}

func LoginHandler(resW http.ResponseWriter, req *http.Request) {
	// Validate request
	if _, ok := req.Header["Email"]; !ok {
		resW.WriteHeader(http.StatusBadRequest)
		resW.Write([]byte("Email Missing."))
		return
	}

	if _, ok := req.Header["PasswordHash"]; !ok {
		resW.WriteHeader(http.StatusBadRequest)
		resW.Write([]byte("PasswordHash Missing."))
		return
	}

	isPassValid, err := validateUser(req.Header["Email"][0], req.Header["PasswordHash"][0])
	if err != nil {
		// User not found
		resW.WriteHeader(http.StatusUnauthorized)
		resW.Write([]byte("User Not Found."))
		return
	}

	if !isPassValid {
		// Password was invalid
		resW.WriteHeader(http.StatusUnauthorized)
		resW.Write([]byte("Invalid Passord."))
		return
	}

	tokenString, err := getSignedToken()
	if err != nil {
		fmt.Println(err)
		resW.WriteHeader(http.StatusInternalServerError)
		resW.Write([]byte("Internal Server Error."))
		return
	}

	resW.WriteHeader(http.StatusOK)
	resW.Write([]byte(tokenString))
}

// All requestes should be authenticated before accessing resource.
func tokenValidater(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resW http.ResponseWriter, req *http.Request) {
		// Check if token is present
		if _, ok := req.Header["Token"]; !ok {
			resW.WriteHeader(http.StatusUnauthorized)
			resW.Write([]byte("Token Missing"))
			return
		}

		token := req.Header["Token"][0]
		isValid, err := jwt.ValidateToken(token, os.Getenv("JWT_PRIVATE_KEY"))
	
		if err != nil {
			resW.WriteHeader(http.StatusInternalServerError)
			resW.Write([]byte("Token Validation Faild."))
			return
		}

		if !isValid {
			resW.WriteHeader(http.StatusUnauthorized)
			resW.Write([]byte("Token Invalid."))
			return
		}
		resW.WriteHeader(http.StatusOK)
		resW.Write([]byte("Authorized"))
	})
}