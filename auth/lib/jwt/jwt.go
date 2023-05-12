package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// Function to generate JWT tokens.

// header, payload -> base64 encoding (In this case, map[string]string)
// signature -> encrypted with SHA256 or other algorithm with private key

// var privateKey = []byte(os.Getenv(""))

func GenerateJWT(header string, payload map[string]string, secret string) (string, error) {
	// Create a new hash, type sha256. We will pass secret to it.
	hash := hmac.New(sha256.New, []byte(secret))

	// Marshal [payload], which is a map. This converts it to JSON string.
	payloadStr, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error generating token!")
		return string(payloadStr), err
	}

	// Have unsigned message ready.
	// Hash it by writing to SHA256.
	unsignedStr := header + string(payloadStr)
	hash.Write([]byte(unsignedStr))
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	// We have token!
	header64 := base64.StdEncoding.EncodeToString([]byte(header))
	payload64 := base64.StdEncoding.EncodeToString(payloadStr)

	tokenStr := header64 + "." + payload64 + "." + signature
	return tokenStr, nil
}


// Validate token
func ValidateJWT(token string, secret string) (bool, error) {
	// strings -> package to manipulate UTF-8 encoded strings

	// JWT has 3 parts separated by '.'
	splitToken := strings.Split(token, ".")

	// If length is not 3, token is corrupt
	if len(splitToken) != 3 {
		return false, nil
	}

	// Decode [header] and [payload] back to strings
	// Create signature by combining them.
	header, err := base64.StdEncoding.DecodeString(splitToken[0])
	if err != nil {
		return false, err
	}

	payload, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		return false, err
	}

	unsignedStr := string(header) + string(payload)
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(unsignedStr))

	// Encode it to get signature 
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	fmt.Println(signature)

	// If both signatures match, token is valid
	if signature != splitToken[2] {
		return false, nil
	}

	return true, nil
}