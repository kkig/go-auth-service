package middleware

import (
	"auth_service/lib/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Check bearer token in request header.
// If bearer token is not present, user is not authenticated.
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := jwt.ValidateJWT(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required."})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}