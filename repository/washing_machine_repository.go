package repository

import (
	"LavanderiaBackend/model"
	"gorm.io/gorm"
)

type WashingMachineRepository struct {
	db *gorm.DB
}

func NewWashingMachineRepository(db *gorm.DB) *WashingMachineRepository {
	return &WashingMachineRepository{db}
}

func (repo *WashingMachineRepository) CreateWashingMachine(WashingMachine *model.WashingMachine) error {
	return repo.db.Create(WashingMachine).Error
}

func (repo *WashingMachineRepository) GetAllWashingMachines() ([]model.WashingMachine, error) {
	var washingMachines []model.WashingMachine
	err := repo.db.Find(&washingMachines).Error
	return washingMachines, err
}

func (repo *WashingMachineRepository) GetWashingMachineByID(id string) (model.WashingMachine, error) {
	var washingMachine model.WashingMachine
	err := repo.db.Where("id = ?", id).First(&washingMachine).Error
	return washingMachine, err
}
