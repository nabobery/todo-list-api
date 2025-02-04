package middlewares

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTAuthMiddleware validates the JWT token and sets the userID in the context.
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		// Expect header format: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		tokenString := parts[1]
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "your_secret_key"
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method.
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}

		// Validate token claims and expiration.
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if exp, ok := claims["exp"].(float64); ok {
				if int64(exp) < time.Now().Unix() {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token expired"})
					return
				}
			}
			// Store the user_id from the token in the Gin context.
			if userID, ok := claims["user_id"].(string); ok {
				c.Set("userID", userID)
			}
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}
	}
}
