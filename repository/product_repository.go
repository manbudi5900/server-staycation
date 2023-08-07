package repository

import (
	"staycation/domain"

	"gorm.io/gorm"
)

type ProductRepositoryContract interface {
	Save(product domain.Product)(domain.Product, error)
	SaveProductImage(productImage domain.ProductImage)(domain.ProductImage, error)
}
type ProductRepository struct {
	db *gorm.DB
  }
func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db:db}
}

func (r ProductRepository) Save(product domain.Product) (domain.Product, error) {

	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
} 
func (r ProductRepository) SaveProductImage(productImage domain.ProductImage) (domain.ProductImage, error) {

	err := r.db.Save(&productImage).Error
	if err != nil {
		return productImage, err
	}
	return productImage, nil
} 