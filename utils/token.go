package utils

import (
	"os"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte(os.Getenv("API_SECRET"))

func GenerateJWTToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		// You can add more custom claims if needed
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
