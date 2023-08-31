package repository

import (
	"staycation/domain"

	"gorm.io/gorm"
)

type ProductRepositoryContract interface {
	Save(product domain.Product)(domain.Product, error)
	Update(product domain.Product)(domain.Product, error)

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
func (r ProductRepository) Update(product domain.Product) (domain.Product, error) {

	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
} 
func (r ProductRepository) SaveProductImage(productImage domain.ProductImage) (domain.ProductImage, error) {
	sqlStatement := `
		INSERT INTO product_images (product_id, url)
		VALUES ($1, $2)`
	err := r.db.Exec(sqlStatement, productImage.HotelID, productImage.Url).Error
	// err := r.db.Save(&productImage).Error
	if err != nil {
		return productImage, err
	}
	return productImage, nil
} 
func (r ProductRepository) SaveProductImageHotel(productImage domain.ProductImage) (domain.ProductImage, error) {
	sqlStatement := `
		INSERT INTO product_images (hotel_id, url)
		VALUES ($1, $2)`
	err := r.db.Exec(sqlStatement, productImage.HotelID, productImage.Url).Error
	// err := r.db.Save(&productImage).Error
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


	
	err := r.db.Preload("ProductImage").Preload("Hotel").Find(&product).Error

	if err != nil {
		return product, err
	}
	
	return product, nil
}
func (r ProductRepository) GetDetailProduct(Slug string) (domain.Product, error){
	
	var product domain.Product
	err := r.db.Preload("ProductImage").Preload("Hotel").Preload("Feature").Preload("Activity").Preload("Category.Products.ProductImage").Preload("Category.Products", "slug != ?",Slug).Where("slug", Slug).First(&product).Error

	if err != nil {
		return product, err
	}
	
	return product, nil
}

func (r ProductRepository) GetCategory() ([]domain.Category, error){
	
	var category []domain.Category
	
	err := r.db.Preload("Products.ProductImage").Find(&category).Error

	if err != nil {
		return category, err
	}
	
	return category, nil
}