package jwt

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"auth_service/lib/data"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Function to generate JWT tokens.

// header, payload -> base64 encoding (In this case, map[string]string)
// signature -> encrypted with SHA256 or other algorithm with private key

// In HS256 algorithm, we sign with private key
// & use public key(usually text-based PEM format) as verification key
// https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/


var secret = []byte(os.Getenv("JWT_PRIVATE_KEY"))

// Login
func GenerateJWT(user data.User) (string, error) {
	// string to int conversion
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))

	// Make JWT token with signing method of ES256 and claims.
	// Claims are attributes.
	// id  - user id
	// aud - audience
	// iss - issuer
	// iat - "issued at". Time token was issued
	// exp - expiration time

	// When you use multiple keys in application, add 'kid' in head of token.
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"iss": "google.com",	// Issuer of token ex. auth server
		"aud": "example.com",	// For idp, this maybe value such as client id	
		"iat": time.Now().Unix(),
		"nonce": "crypto-value",
		"exp": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})

	// Provie shared secret. It should be very complex.
	// This should be passed as Environment variable.
	return  unsignedToken.SignedString(secret)

}

// Subsequent request
func getJWTFromReq(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}

func checkJWTAlg(ctx *gin.Context) (*jwt.Token, error) {
	tokenString := getJWTFromReq(ctx)

	// When you use multiple keys in application, verify 'kid' in header.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		// Validate 'alg' in head of token is what you are expecting
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	return token, err
}

func ValidateJWT(ctx *gin.Context) error {
	token, err := checkJWTAlg(ctx)

	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return nil
	}

	return errors.New("invalid token provided")
}

func LookupUserByUID(ctx *gin.Context) (data.User, error) {
	err := ValidateJWT(ctx)
	if err != nil {
		return data.User{}, err
	}

	token, _ := checkJWTAlg(ctx)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64))

	user, err := data.FindUserById(userId)
	if err != nil {
		return data.User{}, err
	}
	return user, nil
}

// func getSignedJWT(header string, payload map[string]string, secret string) (string, error) {
// 	// Create a new hash, type sha256. We will pass secret to it.
// 	hash := hmac.New(sha256.New, []byte(secret))

// 	// Marshal [payload], which is a map. This converts it to JSON string.
// 	payloadStr, err := json.Marshal(payload)
// 	if err != nil {
// 		fmt.Println("Error generating token!")
// 		return string(payloadStr), err
// 	}

// 	// Have unsigned message ready.
// 	// Hash it by writing to SHA256.
// 	unsignedStr := header + string(payloadStr)
// 	hash.Write([]byte(unsignedStr))
// 	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))

// 	// We have token!
// 	header64 := base64.StdEncoding.EncodeToString([]byte(header))
// 	payload64 := base64.StdEncoding.EncodeToString(payloadStr)

// 	tokenStr := header64 + "." + payload64 + "." + signature
// 	return tokenStr, nil
// }