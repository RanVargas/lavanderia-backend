package repository

import (
	"LavanderiaBackend/model"
	"gorm.io/gorm"
)

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{db}
}

func (repo *ClientRepository) CreateClient(client *model.Client) error {
	return repo.db.Create(client).Error
}

func (repo *ClientRepository) GetAllClients() ([]model.Client, error) {
	var client []model.Client
	err := repo.db.Find(&client).Error
	return client, err
}

func (repo *ClientRepository) GetClientByID(id string) (model.Client, error) {
	var client model.Client
	err := repo.db.Where("id = ?", id).First(&client).Error
	return client, err
}

func (repo *ClientRepository) UpdateClient(client *model.Client) error {
	return repo.db.Save(client).Error
}

func (repo *ClientRepository) DeleteClient(id string) error {
	return repo.db.Where("id = ?", id).Delete(&model.Client{}).Error
}
