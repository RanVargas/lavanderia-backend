package repository

import (
	"LavanderiaBackend/model"
	"github.com/google/uuid"
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

func (repo *WashingMachineRepository) UpdateWashingMachine(washingMachine *model.WashingMachine) error {
	return repo.db.Save(washingMachine).Error
}

func (repo *WashingMachineRepository) DeleteWashingMachine(id string) error {
	return repo.db.Delete(&model.WashingMachine{}, "id = ?", id).Error
}

func (repo *WashingMachineRepository) FetchWashingRequests() ([]model.Request, error) {
	var requests []model.Request
	err := repo.db.Preload("Services").Where("fulfilled = false").Find(&requests).Error
	if err != nil {
		return nil, err
	}

	var washingRequests []model.Request
	for _, req := range requests {
		for _, service := range req.Services {
			if service.Name == "washing" {
				washingRequests = append(washingRequests, req)
				break
			}
		}
	}
	return washingRequests, nil
}

func (repo *WashingMachineRepository) AssignMachineToRequest(machine model.WashingMachine, requestID uuid.UUID) error {
	return repo.db.Model(&machine).Update("current_request_id", requestID).Error
}

func (repo *WashingMachineRepository) GetAvailableMachine() (*model.WashingMachine, error) {
	var machine model.WashingMachine
	result := repo.db.Where("occupied = false").First(&machine)
	return &machine, result.Error
}

func (repo *WashingMachineRepository) SetMachineAvailable(machineId uuid.UUID) error {
	return repo.db.Model(&model.WashingMachine{}).Where("id = ?", machineId).Update("occupied", false).Error
}
