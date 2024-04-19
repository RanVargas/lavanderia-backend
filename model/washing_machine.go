package model

import "gorm.io/gorm"

type WashingMachine struct {
	gorm.Model
	Name  string
	Email string `gorm:"uniqueIndex"`
}
