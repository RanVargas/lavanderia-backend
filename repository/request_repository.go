package repository

import (
	"LavanderiaBackend/model"
	"gorm.io/gorm"
)

type RequestRepository struct {
	db *gorm.DB
}

func NewRequestRepository(db *gorm.DB) *RequestRepository {
	return &RequestRepository{db}
}

func (repo *RequestRepository) CreateRequest(request *model.Request) error {
	return repo.db.Create(request).Error
}

func (repo *RequestRepository) GetAllRequests() ([]model.Request, error) {
	var requests []model.Request
	err := repo.db.Find(&requests).Error
	return requests, err
}

func (repo *RequestRepository) GetRequestByID(id string) (model.Request, error) {
	var request model.Request
	err := repo.db.Where("id = ?", id).First(&request).Error
	return request, err
}

func (repo *RequestRepository) UpdateRequest(request *model.Request) error {
	return repo.db.Save(request).Error
}

func (repo *RequestRepository) DeleteRequestByID(id string) error {
	return repo.db.Where("id = ?", id).Delete(&model.Request{}).Error
}
