package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WashingMachine struct {
	gorm.Model
	MachineModel   string    `json:"machine_model,omitempty"`
	Id             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex;primaryKey" json:"id,omitempty"`
	Capacity       float64   `json:"capacity,omitempty"`
	Occupied       bool      `json:"occupied,omitempty"`
	CurrentRequest *Request  `gorm:"foreignKey:WashingMachineID" json:"current_request,omitempty"`
}
