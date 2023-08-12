package service

import (
	"staycation/domain"
	"staycation/dto"
	"staycation/repository"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)


type ProductServiceContract  interface {
	Save(input dto.ProductInput) (domain.Product, error)

  
  }
  
type ProductService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(m repository.ProductRepository) ProductService{
	return ProductService{m}
}

func (s ProductService) Save(input dto.ProductInput) (domain.Product, error){
	product := domain.Product{}
	product.CategoryID = input.CategoryID
	product.Slug = slug.Make(input.Name)
	product.Name = input.Name
	product.Price = input.Price
	product.Country = input.Country
	product.Province = input.Province
	product.City = input.City
	product.Description = input.Description
	product.ID = uuid.NewString()
	newProduct, err := s.ProductRepository.Save(product)
	if err != nil{
		return newProduct, err
	}
	for i := 0; i < len(input.Url); i++ {
		productImage := domain.ProductImage{}
		productImage.Url = input.Url[i]
		productImage.ProductID = newProduct.ID
		_, err := s.ProductRepository.SaveProductImage(productImage)
		if err != nil{
			return product, err
		}
	}
	return newProduct, nil
}


func (s ProductService) GetLandingPage() (domain.LandingPage, error){
	var landingPage domain.LandingPage
	hero, err := s.ProductRepository.GetHero()
	if err != nil{
		return landingPage, err
	}
	landingPage.Hero = hero
	product, err := s.ProductRepository.GetMostPicked()
	if err != nil{
		return landingPage, err
	}
	landingPage.Product = product
	category, err := s.ProductRepository.GetCategory()
	if err != nil{
		return landingPage, err
	}
	landingPage.Category = category
	return landingPage, nil
	
}
func (s ProductService) GetProductDetail(Slug string) (domain.Product, error){
	var product domain.Product
	product, err := s.ProductRepository.GetDetailProduct(Slug)
	if err != nil{
		return product, err
	}
	
	return product, nil
	
}



