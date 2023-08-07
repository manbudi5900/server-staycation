package service

import (
	"fmt"
	"staycation/domain"
	"staycation/dto"
	"staycation/repository"
)


type ActivityServiceContract interface {
  Save(input dto.ActivityInput) (domain.Activity, error)
  Update(input dto.ActivityInput, id string)(domain.Activity, error)
  GetActivity(limit, page int) ([]domain.Activity, error)

}

type ActivityService struct {
  ActivityRepository repository.ActivityRepository
}

func NewActivityService(m repository.ActivityRepository) ActivityService{
	return ActivityService{m}
}
func (s *ActivityService) Save(input dto.ActivityInput)(domain.Activity, error){
	activity := domain.Activity{}
	activity.Name = input.Name
	activity.Type = input.Type
	activity.ImageUrl = input.ImageUrl

	newActivity, err := s.ActivityRepository.Save(activity)
	if err != nil{
		return newActivity, err
	}
	return newActivity, nil

}
func (s *ActivityService) Update(input dto.ActivityInput, id string)(domain.Activity, error){
	
	
	activity, err := s.ActivityRepository.FindByID(id)
	if err != nil {
		return activity, err
	}
	activity.Name = input.Name
	activity.Type = input.Type
	activity.ImageUrl = input.ImageUrl

	
	fmt.Println(activity)
	updateActivity, err := s.ActivityRepository.Update(activity)
	if err != nil {
		fmt.Println("Errornya disini")
		return activity, err
	}
	return updateActivity, nil

}
func (s *ActivityService) GetActivity(limit, page int)([]domain.Activity, error){
	var Activity []domain.Activity
	Activity, err := s.ActivityRepository.GetActivity(limit, page)
	if err != nil {
		fmt.Println("Errornya disini")
		return Activity, err
	}
	return Activity, nil
}