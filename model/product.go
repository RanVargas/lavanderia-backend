package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model       `json:"gorm_._model"`
	Id               int32   `gorm:"primaryKey" json:"id,omitempty"`
	Name             string  `json:"name,omitempty"`
	Quantity         int32   `json:"quantity,omitempty"`
	RestockThreshold float64 `json:"restock_threshold,omitempty"`
}
