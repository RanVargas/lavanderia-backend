package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"gorm_._model"`
	Id         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex;primaryKey" json:"id,omitempty"`
	Username   string    `gorm:"uniqueIndex;unique" json:"username,omitempty"`
	Email      string    `gorm:"uniqueIndex;unique" json:"email,omitempty"`
	Privileges int       `json:"privileges,omitempty"`
	Password   string    `json:"password,omitempty"`
}
