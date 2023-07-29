package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte(os.Getenv("API_SECRET"))

func JwtMiddleware(c *gin.Context) {
	// Extract the token from the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing access token"})
		c.Abort()
		return
	}

	// Check if the token starts with "Bearer "
	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid access token format"})
		c.Abort()
		return
	}

	// Extract the token from the header
	tokenString := authHeader[len("Bearer "):]

	// Verify the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid access token"})
		c.Abort()
		return
	}

	// Pass the email from the token to the context
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		email, _ := claims["email"].(string)
		c.Set("email", email)
	}

	c.Next()
}
