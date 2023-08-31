package repository

import (
	"staycation/domain"

	"gorm.io/gorm"
)

type HotelRepositoryContract interface {
	Gethotel(limit, page int) ([]domain.Hotel, error)
	Save(hotel domain.Hotel) (domain.Hotel, error)
	FindByID(ID string) (domain.Hotel, error)
	Update(hotel domain.Hotel)(domain.Hotel, error)
}
type HotelRepository struct {
	db *gorm.DB
  }
func NewHotelRepository(db *gorm.DB) HotelRepository {
	return HotelRepository{db:db}
}

func (r HotelRepository) GetHotel(limit, page int) ([]domain.Hotel, error) {
	var hotel []domain.Hotel

	offset := (page - 1) * limit

	err := r.db.Offset(offset).
	Limit(limit).Find(&hotel).Error

	if err != nil {
		return hotel, err
	}
	return hotel, nil
}

func (r HotelRepository) Save(hotel domain.Hotel) (domain.Hotel, error) {
	err := r.db.Create(&hotel).Error
	if err != nil {
		return hotel, err
	}
	return hotel, nil
}
func (r HotelRepository) Update(hotel domain.Hotel)(domain.Hotel, error){
	err := r.db.Save(&hotel).Error
	if err != nil {
		return hotel, err
	}
	return hotel, nil
}
func (r HotelRepository) FindByID(ID string) (domain.Hotel, error) {
	var hotel domain.Hotel
	err := r.db.Where("id =?", ID).Find(&hotel).Error
	if err != nil {
		return hotel, err
	}
	return hotel, nil
}
func (r HotelRepository) FindByIDDetail(ID string) (domain.Hotel, error) {
	var hotel domain.Hotel
	err := r.db.Where("id =?", ID).Preload("ProductImageHotel").Preload("Feature").Find(&hotel).Error
	if err != nil {
		return hotel, err
	}
	return hotel, nil
}

