package auth

import (
	"LavanderiaBackend/model"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("my_secret_key") // Should be in a secure place or from environment variables

type Claims struct {
	User model.User `json:"user"`
	jwt.StandardClaims
}

// GenerateToken generates a jwt token and assign a username to it's claims and return it
func GenerateToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour) // Token expires after 1 hour
	claims := &Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	return tokenString, err
}

/*func ValidateToken(signedToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || claims.User.Id == uuid.Nil { // Added check to ensure user ID is not nil
		return nil, fmt.Errorf("couldn't parse claims or invalid user")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, fmt.Errorf("jwt is expired")
	}

	return claims, nil
}*/

func ValidateToken(signedToken string) (*Claims, error) {
	// Parse the token.
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	// Check if the token's claims can be used and are valid.
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("could not validate claims")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, fmt.Errorf("jwt is expired")
	}

	return claims, nil
}
