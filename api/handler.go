package api

import (
	"LavanderiaBackend/model"
	"LavanderiaBackend/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context, repo *repository.UserRepository) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := repo.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func GetAllUsers(c *gin.Context, repo *repository.UserRepository) {
	users, err := repo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context, repo *repository.UserRepository) {
	id := c.Param("id")
	user, err := repo.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateProduct(c *gin.Context, repo *repository.ProductRepository) {
	var product model.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := repo.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"product": product})
}

func GetAllProducts(c *gin.Context, repo *repository.ProductRepository) {
	products, err := repo.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func GetProductByName(c *gin.Context, repo *repository.ProductRepository) {
	id := c.Param("name")
	product, err := repo.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func CreateClient(c *gin.Context, repo *repository.ClientRepository) {
	var client model.Client
	if err := c.BindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := repo.CreateClient(&client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"client": client})
}

func GetAllClients(c *gin.Context, repo *repository.ClientRepository) {
	clients, err := repo.GetAllClients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clients)
}

func GetClientByID(c *gin.Context, repo *repository.ClientRepository) {
	id := c.Param("id")
	client, err := repo.GetClientByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, client)
}

func CreateRequest(c *gin.Context, repo *repository.RequestRepository) {
	var request model.Request
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := repo.CreateRequest(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"product": request})
}

func GetAllRequests(c *gin.Context, repo *repository.RequestRepository) {
	requests, err := repo.GetAllRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, requests)
}

func GetRequestByID(c *gin.Context, repo *repository.RequestRepository) {
	id := c.Param("id")
	request, err := repo.GetRequestByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, request)
}

func CreateService(c *gin.Context, repo *repository.ServiceRepository) {
	var service model.Service
	if err := c.BindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := repo.CreateService(&service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"service": service})
}

func GetAllServices(c *gin.Context, repo *repository.ServiceRepository) {
	services, err := repo.GetAllServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

func GetServiceByID(c *gin.Context, repo *repository.ServiceRepository) {
	id := c.Param("id")
	service, err := repo.GetServiceByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service)
}

func CreateWashingMachine(c *gin.Context, repo *repository.WashingMachineRepository) {
	var washingMachine model.WashingMachine
	if err := c.BindJSON(&washingMachine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := repo.CreateWashingMachine(&washingMachine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"washingMachine": washingMachine})
}

func GetAllWashingMachines(c *gin.Context, repo *repository.WashingMachineRepository) {
	washingMachines, err := repo.GetAllWashingMachines()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, washingMachines)
}

func GetWashingMachineByID(c *gin.Context, repo *repository.WashingMachineRepository) {
	id := c.Param("id")
	washingMachine, err := repo.GetWashingMachineByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, washingMachine)
}
