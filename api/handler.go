package api

import (
	"LavanderiaBackend/api/auth"
	"LavanderiaBackend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func AddExampleRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/examples", func(c *gin.Context) {
		var examples []model.User
		if result := db.Find(&examples); result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, examples)
	})

	r.POST("/examples", func(c *gin.Context) {
		var example model.User
		if err := c.ShouldBindJSON(&example); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if result := db.Create(&example); result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, example)
	})
}

func LoginForUsers(r *gin.Engine) {
	r.POST("/login", func(c *gin.Context) {
		// Assume you have some way of validating credentials
		username := c.PostForm("username")
		// Simulating credential validation
		if username != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "credentials are invalid"})
			return
		}

		token, err := auth.GenerateToken(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})
}
