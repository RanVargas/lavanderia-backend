package main

import (
	"LavanderiaBackend/api"
	"LavanderiaBackend/config"
	"LavanderiaBackend/repository"
	services "LavanderiaBackend/service"
	"github.com/gin-gonic/gin"
	"log"
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
	service := services.NewAssignmentService(washingMachineRepo)

	r := gin.Default()
	r.Use(gin.Logger())

	go service.StartAssignmentProcess()

	authGroup := r.Group("/")
	{
		authGroup.POST("/register", func(c *gin.Context) {
			api.RegisterUser(c, userRepo)
		})
		authGroup.POST("/login", func(c *gin.Context) {
			api.LoginForUsers(c, userRepo)
		})

		// Protected routes
		authGroup.Use(api.AuthMiddleware(userRepo))
		{
			// Users routes
			authGroup.POST("/users", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.CreateUser(c, userRepo)
			})
			authGroup.GET("/users", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.GetAllUsers(c, userRepo)
			})
			authGroup.GET("/users/:id", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.GetUserByID(c, userRepo)
			})
			authGroup.PATCH("/users/:id", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.UpdateUser(c, userRepo)
			})
			authGroup.DELETE("/users/:id", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.DeleteUser(c, userRepo)
			})

			// Product routes
			authGroup.POST("/products", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.CreateProduct(c, productRepo)
			})
			authGroup.GET("/products", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.GetAllProducts(c, productRepo)
			})
			authGroup.GET("/products/:name", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.GetProductByName(c, productRepo)
			})
			authGroup.PATCH("/products/:name", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.UpdateProduct(c, productRepo)
			})
			authGroup.DELETE("/products/:name", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.DeleteProduct(c, productRepo)
			})

			// Requests routes
			authGroup.POST("/requests", api.PrivilegeMiddleware(2), func(c *gin.Context) {
				api.CreateRequest(c, requestRepo)
			})
			authGroup.GET("/requests", api.PrivilegeMiddleware(2), func(c *gin.Context) {
				api.GetAllRequests(c, requestRepo)
			})
			authGroup.GET("/requests/:id", api.PrivilegeMiddleware(2), func(c *gin.Context) {
				api.GetRequestByID(c, requestRepo)
			})
			authGroup.PATCH("/requests/:id", api.PrivilegeMiddleware(2), func(c *gin.Context) {
				api.UpdateRequest(c, requestRepo)
			})
			authGroup.DELETE("/requests/:id", api.PrivilegeMiddleware(2), func(c *gin.Context) {
				api.DeleteRequest(c, requestRepo)
			})

			// Clients routes
			authGroup.POST("/clients", api.PrivilegeMiddleware(1), func(c *gin.Context) {
				api.CreateClient(c, clientRepo)
			})
			authGroup.GET("/clients", api.PrivilegeMiddleware(1), func(c *gin.Context) {
				api.GetAllClients(c, clientRepo)
			})
			authGroup.GET("/clients/:id", api.PrivilegeMiddleware(1), func(c *gin.Context) {
				api.GetClientByID(c, clientRepo)
			})
			authGroup.PATCH("/clients/:id", api.PrivilegeMiddleware(1), func(c *gin.Context) {
				api.UpdateClient(c, clientRepo)
			})
			authGroup.DELETE("/clients/:id", api.PrivilegeMiddleware(1), func(c *gin.Context) {
				api.DeleteClient(c, clientRepo)
			})

			// washingMachines routes
			authGroup.POST("/washingMachines", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.CreateWashingMachine(c, washingMachineRepo)
			})
			authGroup.GET("/washingMachines", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.GetAllWashingMachines(c, washingMachineRepo)
			})
			authGroup.GET("/washingMachines/:id", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.GetWashingMachineByID(c, washingMachineRepo)
			})
			authGroup.PATCH("/washingMachines/:id", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.UpdateWashingMachine(c, washingMachineRepo)
			})
			authGroup.DELETE("/washingMachines/:id", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.DeleteWashingMachine(c, washingMachineRepo)
			})

			// Services routes
			authGroup.POST("/services", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.CreateService(c, serviceRepo)
			})
			authGroup.GET("/services", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.GetAllServices(c, serviceRepo)
			})
			authGroup.GET("/services/:id", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.GetServiceByID(c, serviceRepo)
			})
			authGroup.PATCH("/services/:id", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.UpdateService(c, serviceRepo)
			})
			authGroup.DELETE("/services/:id", api.PrivilegeMiddleware(0), func(c *gin.Context) {
				api.DeleteService(c, serviceRepo)
			})
		}
	}

	if err := r.Run(":7575"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
