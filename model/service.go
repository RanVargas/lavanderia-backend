package model

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Id    int     `gorm:"primaryKey" json:"id"`
	Name  string  `gorm:"unique" json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}
