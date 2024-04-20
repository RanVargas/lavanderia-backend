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

	db, err := repository.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	productRepo := repository.NewProductRepository(db)
	userRepo := repository.NewUserRepository(db)
	serviceRepo := repository.NewServiceRepository(db)
	washingMachineRepo := repository.NewWashingMachineRepository(db)
	requestRepo := repository.NewRequestRepository(db)
	clientRepo := repository.NewClientRepository(db)

	r := gin.Default()
	r.Use(gin.Logger())

	authGroup := r.Group("/")
	{
		authGroup.POST("/register", func(c *gin.Context) {
			api.RegisterUser(c, userRepo)
		})
		authGroup.POST("/login", func(c *gin.Context) {
			api.LoginForUsers(c, userRepo)
		})

		// Protected routes
		authGroup.Use(api.AuthMiddleware())
		{
			// Users routes
			authGroup.POST("/users", func(c *gin.Context) {
				api.CreateUser(c, userRepo)
			})
			authGroup.GET("/users", func(c *gin.Context) {
				api.GetAllUsers(c, userRepo)
			})
			authGroup.GET("/users/:id", func(c *gin.Context) {
				api.GetUserByID(c, userRepo)
			})

			// Product routes
			authGroup.POST("/products", func(c *gin.Context) {
				api.CreateProduct(c, productRepo)
			})
			authGroup.GET("/products", func(c *gin.Context) {
				api.GetAllProducts(c, productRepo)
			})
			authGroup.GET("/products/:name", func(c *gin.Context) {
				api.GetProductByName(c, productRepo)
			})

			// Requests routes
			authGroup.POST("/requests", func(c *gin.Context) {
				api.CreateRequest(c, requestRepo)
			})
			authGroup.GET("/requests", func(c *gin.Context) {
				api.GetAllRequests(c, requestRepo)
			})
			authGroup.GET("/requests/:id", func(c *gin.Context) {
				api.GetRequestByID(c, requestRepo)
			})

			// Clients routes
			authGroup.POST("/clients", func(c *gin.Context) {
				api.CreateClient(c, clientRepo)
			})
			authGroup.GET("/clients", func(c *gin.Context) {
				api.GetAllClients(c, clientRepo)
			})
			authGroup.GET("/clients/:id", func(c *gin.Context) {
				api.GetClientByID(c, clientRepo)
			})

			// washingMachines routes
			authGroup.POST("/washingMachines", func(c *gin.Context) {
				api.CreateWashingMachine(c, washingMachineRepo)
			})
			authGroup.GET("/washingMachines", func(c *gin.Context) {
				api.GetAllWashingMachines(c, washingMachineRepo)
			})
			authGroup.GET("/washingMachines/:id", func(c *gin.Context) {
				api.GetWashingMachineByID(c, washingMachineRepo)
			})

			// Services routes
			authGroup.POST("/services", func(c *gin.Context) {
				api.CreateService(c, serviceRepo)
			})
			authGroup.GET("/services", func(c *gin.Context) {
				api.GetAllServices(c, serviceRepo)
			})
			authGroup.GET("/services/:id", func(c *gin.Context) {
				api.GetServiceByID(c, serviceRepo)
			})
		}
	}

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
