package model

import "gorm.io/gorm"

type Service struct {
	gorm.Model  `json:"gorm_._model"`
	Id          int        `gorm:"primaryKey" json:"id" json:"id,omitempty"`
	Name        string     `gorm:"unique" json:"name,omitempty" json:"name,omitempty"`
	Price       float64    `json:"price,omitempty" json:"price,omitempty"`
	IsWashing   bool       `gorm:"default:false" json:"isWashing,omitempty"`
	IsDrying    bool       `gorm:"default:false" json:"isDrying,omitempty"`
	IsFullCycle bool       `gorm:"default:true" json:"isFullCycle,omitempty"`
	Products    []*Product `gorm:"many2many:service_products;" json:"products,omitempty" json:"products,omitempty"`
}
