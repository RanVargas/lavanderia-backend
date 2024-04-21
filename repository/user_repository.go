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

func (repo *UserRepository) GetUserByID(id string) (model.User, error) {
	var user model.User
	err := repo.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (repo *UserRepository) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	result := repo.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func (repo *UserRepository) UpdateUser(user *model.User) error {
	return repo.db.Save(user).Error
}

func (repo *UserRepository) DeleteUser(id string) error {
	return repo.db.Where("id = ?", id).Delete(&model.User{}).Error
}
