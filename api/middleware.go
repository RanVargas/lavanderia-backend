package api

import (
	"LavanderiaBackend/model"
	"LavanderiaBackend/repository"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"

	"LavanderiaBackend/api/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(userRepo *repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token not found"})
			return
		}

		claims, err := auth.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error() + " At validating claims"})
			return
		}

		// Fetch the user from the database
		user, err := userRepo.GetUserByUsername(claims.User.Username)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		// Store the user model in the context
		c.Set("user", user)

		c.Next()
	}
}

/*func PrivilegeMiddleware(requiredPrivilege int) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
			return
		}

		userModel := user.(model.User)
		if userModel.Privileges > requiredPrivilege {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient privileges"})
			return
		}

		c.Next()
	}
}*/

func PrivilegeMiddleware(requiredPrivilege int) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
			return
		}

		userModel := user.(model.User)
		if userModel.Privileges <= requiredPrivilege {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient privileges"})
		}
	}
}

func RegisterUser(c *gin.Context, userRepo *repository.UserRepository) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: " + err.Error()})
		return
	}

	// Here, you could add hashing for the password before saving
	var err error
	user.Password, err = auth.HashingPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	err = userRepo.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user: " + err.Error()})
		return
	}

	// Create JWT token after successful registration
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &auth.Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(auth.JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": tokenString, "user": user, "username": user.Username})
}

func CheckUserCredentials(username, password string, repo *repository.UserRepository) bool {
	// Fetch user from the database by username
	user, err := repo.GetUserByUsername(username)
	if err != nil {
		return false
	}

	return auth.CheckPasswordHash(password, user.Password)
}

func LoginForUsers(c *gin.Context, userRepo *repository.UserRepository) {
	/*var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}*/
	var credentials model.User

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	if !CheckUserCredentials(credentials.Username, credentials.Password, userRepo) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &auth.Claims{
		User: credentials,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(auth.JwtKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
