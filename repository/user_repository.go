package repository

import (
	"LavanderiaBackend/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) CreateUser(user *model.User) error {
	return repo.db.Create(user).Error
}

func (repo *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := repo.db.Find(&users).Error
	return users, err
}
