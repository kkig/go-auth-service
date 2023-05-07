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

// encrypt the header and payload together using a secret key to get signature
// Validate token, decode header and payload. Then combine them to create new signature
// If this matches the signature present in the token, token is valid.

// var privateKey = []byte(os.Getenv(""))

func GenerateToken(header string, payload map[string]string, secret string) (string, error) {
	// Create a new hash, type sha256. We will pass secret to it.
	hash := hmac.New(sha256.New, []byte(secret))
	header64 := base64.StdEncoding.EncodeToString([]byte(header))

	// Marshal payload, which is a map. This converts it to JSON string.
	payloadstr, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error generating token!")
		return string(payloadstr), err
	}

	payload64 := base64.StdEncoding.EncodeToString(payloadstr)

	// Now add encoded string.
	message := header64 + "." + payload64

	// Have unsigned message ready.
	unsignedStr := header + string(payloadstr)

	// Write to SHA256, to hash it.
	hash.Write([]byte(unsignedStr))
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	// We have token!
	tokenStr := message + "." + signature
	return tokenStr, nil
}


// This helps in validating token
func ValidateToken(token string, secret string) (bool, error) {
	// strings -> package to manipulate UTF-8 encoded strings

	// JWT has 3 parts separated by '.'
	splitToken := strings.Split(token, ".")

	// If length is not 3, token is corrupt
	if len(splitToken) != 3 {
		return false, nil
	}

	// Decode header and payload back to strings
	header, err := base64.StdEncoding.DecodeString(splitToken[0])
	if err != nil {
		return false, err
	}

	payload, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		return false, err
	}

	// Create signature
	unsignedStr := string(header) + string(payload)
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(unsignedStr))

	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	fmt.Println(signature)

	// If both signatures don't match, token is incorrect
	if signature != splitToken[2] {
		return false, nil
	}

	return true, nil
}