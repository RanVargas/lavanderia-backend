package main

import (
	"LavanderiaBackend/api"
	"LavanderiaBackend/config"
	"LavanderiaBackend/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	_, err = repository.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := gin.Default()
	r.Use(gin.Logger())

	authGroup := r.Group("/")
	authGroup.Use(api.AuthMiddleware())
	{
		//authGroup.GET("/examples", exampleHandler)
		authGroup.POST("/login", api.LoginForUsers)
	}

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
