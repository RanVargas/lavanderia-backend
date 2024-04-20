package repository

import (
	"LavanderiaBackend/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (repo *ProductRepository) CreateProduct(product *model.Product) error {
	return repo.db.Create(product).Error
}

func (repo *ProductRepository) GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	err := repo.db.Find(&products).Error
	return products, err
}

func (repo *ProductRepository) GetProductByID(id string) (model.Product, error) {
	var product model.Product
	err := repo.db.Where("name = ?", id).First(&product).Error
	return product, err
}

func (repo *ProductRepository) UpdateProduct(product *model.Product) error {
	return repo.db.Save(product).Error
}

func (repo *ProductRepository) DeleteProduct(id string) error {
	return repo.db.Where("id = ?", id).Delete(&model.Product{}).Error
}
