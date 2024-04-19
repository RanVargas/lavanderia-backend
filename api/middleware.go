package api

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement your authentication logic here
		// If unauthorized, set an appropriate response and return
		c.Next()
	}
}
