package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model       `json:"gorm_._model"`
	Name             string  `gorm:"uniqueIndex;unique;primaryKey" json:"name,omitempty"`
	Quantity         int32   `json:"quantity,omitempty"`
	RestockThreshold float64 `json:"restock_threshold,omitempty"`
}
