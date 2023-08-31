package repository

import (
	"staycation/domain"

	"gorm.io/gorm"
)

type FeatureRepositoryContract interface {
	GetFeature(limit, page int) ([]domain.Feature, error)
	Save(feature domain.Feature) (domain.Feature, error)
	FindByID(ID string) (domain.Feature, error)
	Update(feature domain.Feature)(domain.Feature, error)
}
type FeatureRepository struct {
	db *gorm.DB
  }
func NewFeatureRepository(db *gorm.DB) FeatureRepository {
	return FeatureRepository{db:db}
}

func (r FeatureRepository) GetFeature(limit, page int) ([]domain.Feature, error) {
	var feature []domain.Feature

	offset := (page - 1) * limit

	err := r.db.Offset(offset).
	Limit(limit).Find(&feature).Error

	if err != nil {
		return feature, err
	}
	return feature, nil
}

func (r FeatureRepository) Save(feature domain.Feature) (domain.Feature, error) {
	err := r.db.Create(&feature).Error
	if err != nil {
		return feature, err
	}
	return feature, nil
}
func (r FeatureRepository) SaveHotel(feature domain.Feature) (domain.Feature, error) {
	sqlStatement := `
		INSERT INTO features (qty, name, image_url, hotel_id)
		VALUES ($1, $2, $3, $4)`
	err := r.db.Exec(sqlStatement, feature.Qty, feature.Name, feature.ImageUrl, feature.HotelID).Error
	// err := r.db.Create(&feature).Error
	if err != nil {
		return feature, err
	}
	return feature, nil
}
func (r FeatureRepository) Update(feature domain.Feature)(domain.Feature, error){
	err := r.db.Save(&feature).Error
	if err != nil {
		return feature, err
	}
	return feature, nil
}
func (r FeatureRepository) FindByID(ID string) (domain.Feature, error) {
	var feature domain.Feature
	err := r.db.Where("id =?", ID).Find(&feature).Error
	if err != nil {
		return feature, err
	}
	return feature, nil
}

