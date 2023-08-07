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
func (r ProductRepository) GetHero() (domain.Hero, error){
	var feature []domain.Feature
	var hero domain.Hero
	var activity []domain.Activity
	var product []domain.Product


	err := r.db.Find(&feature).Error

	if err != nil {
		return hero, err
	}
	err = r.db.Find(&activity).Error

	if err != nil {
		return hero, err
	}
	err = r.db.Find(&product).Error

	if err != nil {
		return hero, err
	}
	hero.Cities = len(product)
	hero.Travelers = len(activity)
	hero.Treasures = len(feature)
	return hero, nil
}

func (r ProductRepository) GetMostPicked() ([]domain.Product, error){
	
	var product []domain.Product


	
	err := r.db.Preload("ProductImage").Find(&product).Error

	if err != nil {
		return product, err
	}
	
	return product, nil
}
func (r ProductRepository) GetCategory() ([]domain.Category, error){
	
	var category []domain.Category


	
	err := r.db.Preload("Product.ProductImage").Find(&category).Error

	if err != nil {
		return category, err
	}
	
	return category, nil
}