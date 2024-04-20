package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("my_secret_key") // Should be in a secure place or from environment variables

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken generates a jwt token and assign a username to it's claims and return it
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour) // Token expires after 1 hour
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	return tokenString, err
}

// ValidateToken validates the jwt token
func ValidateToken(signedToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, fmt.Errorf("jwt is expired")
	}

	return claims, nil
}
