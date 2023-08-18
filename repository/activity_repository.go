package repository

import (
	"staycation/domain"

	"gorm.io/gorm"
)

type ActivityRepositoryContract interface {
	GetActivity(limit, page int) ([]domain.Activity, error)
	Save(Activity domain.Activity) (domain.Activity, error)
	FindByID(ID string) (domain.Activity, error)
	Update(Activity domain.Activity)(domain.Activity, error)
}
type ActivityRepository struct {
	db *gorm.DB
  }
func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return ActivityRepository{db:db}
}

func (r ActivityRepository) GetActivity(limit, page int) ([]domain.Activity, error) {
	var activity []domain.Activity

	offset := (page - 1) * limit

	err := r.db.Offset(offset).
	Limit(limit).Find(&activity).Error

	if err != nil {
		return activity, err
	}
	return activity, nil
}

func (r ActivityRepository) Save(activity domain.Activity) (domain.Activity, error) {
	err := r.db.Create(&activity).Error
	if err != nil {
		return activity, err
	}
	return activity, nil
}
func (r ActivityRepository) Update(activity domain.Activity)(domain.Activity, error){
	err := r.db.Save(&activity).Error
	if err != nil {
		return activity, err
	}
	return activity, nil
}
func (r ActivityRepository) FindByID(ID string) (domain.Activity, error) {
	var activity domain.Activity
	err := r.db.Where("id =?", ID).Find(&activity).Error
	if err != nil {
		return activity, err
	}
	return activity, nil
}

