package service

import (
	"fmt"
	"staycation/domain"
	"staycation/dto"
	"staycation/repository"
)


type FeatureServiceContract interface {
  Save(input dto.FeatureInput) (domain.Feature, error)
  Update(input dto.FeatureInput, id string)(domain.Feature, error)
  GetFeature(limit, page int) ([]domain.Feature, error)

}

type FeatureService struct {
  FeatureRepository repository.FeatureRepository
}

func NewFeatureService(m repository.FeatureRepository) FeatureService{
	return FeatureService{m}
}
func (s *FeatureService) Save(input dto.FeatureInput)(domain.Feature, error){
	Feature := domain.Feature{}
	Feature.Name = input.Name
	Feature.Qty = input.Qty
	Feature.ImageUrl = input.ImageUrl

	newFeature, err := s.FeatureRepository.Save(Feature)
	if err != nil{
		return newFeature, err
	}
	return newFeature, nil

}
func (s *FeatureService) Update(input dto.FeatureInput, id string)(domain.Feature, error){
	
	
	Feature, err := s.FeatureRepository.FindByID(id)
	if err != nil {
		return Feature, err
	}
	Feature.Name = input.Name
	Feature.Qty = input.Qty
	Feature.ImageUrl = input.ImageUrl

	
	fmt.Println(Feature)
	updateFeature, err := s.FeatureRepository.Update(Feature)
	if err != nil {
		fmt.Println("Errornya disini")
		return Feature, err
	}
	return updateFeature, nil

}
func (s *FeatureService) GetFeature(limit, page int)([]domain.Feature, error){
	var Feature []domain.Feature
	Feature, err := s.FeatureRepository.GetFeature(limit, page)
	if err != nil {
		fmt.Println("Errornya disini")
		return Feature, err
	}
	return Feature, nil
}