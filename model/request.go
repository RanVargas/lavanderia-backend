package model

import "gorm.io/gorm"

type Request struct {
	gorm.Model
	Name  string
	Email string `gorm:"uniqueIndex"`
}
