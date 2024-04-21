package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Request struct {
	gorm.Model
	Id               uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;uniqueIndex" json:"id"`
	OrderedDate      time.Time       `gorm:"uniqueIndex" json:"ordered_date"`
	FulfilledDate    time.Time       `gorm:"uniqueIndex" json:"fulfilled_date"`
	Services         []*Service      `gorm:"many2many:request_services;" json:"services"`
	Fulfilled        bool            `json:"fulfilled"`
	Ongoing          bool            `json:"ongoing"`
	WashingMachineID *uuid.UUID      `gorm:"type:uuid;default:null" json:"washing_machine_id,omitempty"`   // Pointer to allow null
	WashingMachine   *WashingMachine `gorm:"foreignKey:WashingMachineID" json:"washing_machine,omitempty"` // Pointer to allow null
	GrandTotal       float64         `json:"grand_total"`
	Client           Client          `gorm:"foreignKey:ClientID" json:"client"`
	ClientID         uuid.UUID       `gorm:"type:uuid" json:"client_id"`
}

func (r *Request) RequiresWashing() bool {
	for _, service := range r.Services {
		if service.IsWashing || service.IsDrying || service.IsFullCycle {
			return true
		}
	}
	return false
}
