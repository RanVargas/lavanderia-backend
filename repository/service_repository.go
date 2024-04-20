package repository

import (
	"LavanderiaBackend/model"
	"gorm.io/gorm"
)

type ServiceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{db}
}

func (repo *ServiceRepository) CreateService(service *model.Service) error {
	return repo.db.Create(service).Error
}

func (repo *ServiceRepository) GetAllServices() ([]model.Service, error) {
	var services []model.Service
	err := repo.db.Find(&services).Error
	return services, err
}

func (repo *ServiceRepository) GetServiceByID(id string) (model.Service, error) {
	var service model.Service
	err := repo.db.Where("id = ?", id).First(&service).Error
	return service, err
}
