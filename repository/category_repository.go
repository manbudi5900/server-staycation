package repository

import (
	"staycation/domain"

	"gorm.io/gorm"
)

type CategoryRepositoryContract interface {
	GetCategory(limit, page int) ([]domain.Category, error)
	Save(category domain.Category) (domain.Category, error)
	FindByID(ID string) (domain.Category, error)
	Update(category domain.Category)(domain.Category, error)
}
type CategoryRepository struct {
	db *gorm.DB
  }
func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return CategoryRepository{db:db}
}

func (r CategoryRepository) GetCategory(limit, page int) ([]domain.Category, error) {
	var category []domain.Category

	offset := (page - 1) * limit

	err := r.db.Offset(offset).
	Limit(limit).Find(&category).Error

	if err != nil {
		return category, err
	}
	return category, nil
}

func (r CategoryRepository) Save(category domain.Category) (domain.Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}
func (r CategoryRepository) Update(category domain.Category)(domain.Category, error){
	err := r.db.Save(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}
func (r CategoryRepository) FindByID(ID string) (domain.Category, error) {
	var category domain.Category
	err := r.db.Where("id =?", ID).Find(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

