package api

import (
	"LavanderiaBackend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddExampleRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/examples", func(c *gin.Context) {
		var examples []model.ExampleModel
		if result := db.Find(&examples); result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, examples)
	})

	r.POST("/examples", func(c *gin.Context) {
		var example model.ExampleModel
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
