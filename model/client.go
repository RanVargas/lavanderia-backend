package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model `json:"gorm_._model"`
	Id         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex;primaryKey" json:"id,omitempty"`
	Name       string    `json:"username,omitempty"`
	Email      string    `json:"email,omitempty"`
	Address    string    `json:"address,omitempty"`
}
