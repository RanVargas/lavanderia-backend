package model

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &WashingMachine{}, &Client{}, &Request{}, &Service{}, &Product{})
}
