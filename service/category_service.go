package service

import (
	"fmt"
	"staycation/domain"
	"staycation/dto"
	"staycation/repository"
)


type CategoryServiceContract interface {
  Save(input dto.CategoryInput) (domain.Category, error)
  Update(input dto.CategoryInput, id string)(domain.Category, error)
  GetCategory(limit, page int) ([]domain.Category, error)

}

type CategoryService struct {
  CategoryRepository repository.CategoryRepository
}

func NewCategoryService(m repository.CategoryRepository) CategoryService{
	return CategoryService{m}
}
func (s *CategoryService) Save(input dto.CategoryInput)(domain.Category, error){
	category := domain.Category{}
	category.Name = input.Name

	newCategory, err := s.CategoryRepository.Save(category)
	if err != nil{
		return newCategory, err
	}
	return newCategory, nil

}
func (s *CategoryService) Update(input dto.CategoryInput, id string)(domain.Category, error){
	
	
	category, err := s.CategoryRepository.FindByID(id)
	if err != nil {
		return category, err
	}
	category.Name = input.Name

	
	fmt.Println(category)
	updateCategory, err := s.CategoryRepository.Update(category)
	if err != nil {
		fmt.Println("Errornya disini")
		return category, err
	}
	return updateCategory, nil

}
func (s *CategoryService) GetCategory(limit, page int)([]domain.Category, error){
	var category []domain.Category
	category, err := s.CategoryRepository.GetCategory(limit, page)
	if err != nil {
		fmt.Println("Errornya disini")
		return category, err
	}
	return category, nil
}