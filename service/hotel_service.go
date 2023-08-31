package service

import (
	"fmt"
	"staycation/domain"
	"staycation/dto"
	"staycation/repository"
	"strconv"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)


type HotelServiceContract interface {
  Save(input dto.HotelInput) (domain.Hotel, error)
  Update(input dto.HotelInput, id string)(domain.Hotel, error)
  GetHotel(limit, page int) ([]domain.Hotel, error)

}

type HotelService struct {
  HotelRepository repository.HotelRepository
   ProductRepository repository.ProductRepository
   FeatureRepository repository.FeatureRepository
}

func NewHotelService(m repository.HotelRepository, p repository.ProductRepository, f repository.FeatureRepository) HotelService{
	return HotelService{m,p,f}
}
func (s *HotelService) GetHotel(limit, page int)([]domain.Hotel, error){
	var Hotel []domain.Hotel
	Hotel, err := s.HotelRepository.GetHotel(limit, page)
	if err != nil {
		return Hotel, err
	}
	return Hotel, nil
}
func (s *HotelService) Save(input dto.HotelInput)(domain.Hotel, error){
	hotel := domain.Hotel{}
	hotel.ID = uuid.NewString()
	hotel.Name = input.Name
	hotel.Desc = input.Desc
	hotel.Logo = input.Logo
	hotel.Province = input.Province
	hotel.City = input.City 
	hotel.Slug = slug.Make(input.Name)

	_, err := s.HotelRepository.Save(hotel)
	if err != nil {
		fmt.Println("Errornya disini")
		return hotel, err
	}

	


	for i := 0; i < len(input.Images); i++ {
		productImage := domain.ProductImage{}
		productImage.Url = input.Images[i]
		productImage.HotelID = hotel.ID
		_, err := s.ProductRepository.SaveProductImageHotel(productImage)
		if err != nil{
			return hotel, err
		}
	}
	for i := 0; i < len(input.Features); i++ {
		featureHotel := domain.Feature{}
		featureHotel.HotelID = hotel.ID
		featureHotel.Name = input.Features[i][0]
		j, err := strconv.Atoi(input.Features[i][1])
		if err != nil {
			return hotel, err
		}

		featureHotel.Qty = j
		featureHotel.ImageUrl = input.Features[i][2]
		featureHotel.HotelID = hotel.ID
		_, err = s.FeatureRepository.SaveHotel(featureHotel)
		if err != nil{
			return hotel, err
		}
	}

	
	
	hotel, err = s.HotelRepository.FindByIDDetail(hotel.ID)
	if err != nil {
		return hotel, err
	}
	return hotel, nil

}
func (s *HotelService) Update(input dto.HotelInput, id string)(domain.Hotel, error){
	hotel, err := s.HotelRepository.FindByID(id)
	if err != nil {
		return hotel, err
	}
	hotel.Name = input.Name
	hotel.Desc = input.Desc
	hotel.Logo = input.Logo
	hotel.Province = input.Province
	hotel.City = input.City 
	hotel.Slug = slug.Make(input.Name)

	


	for i := 0; i < len(input.Images); i++ {
		fmt.Println("ini dia masuk")
		productImage := domain.ProductImage{}
		productImage.Url = input.Images[i]
		productImage.HotelID = hotel.ID
		_, err := s.ProductRepository.SaveProductImageHotel(productImage)
		if err != nil{
			return hotel, err
		}
	}
	for i := 0; i < len(input.Features); i++ {
		featureHotel := domain.Feature{}
		featureHotel.HotelID = hotel.ID
		featureHotel.Name = input.Features[i][0]
		j, err := strconv.Atoi(input.Features[i][1])
		if err != nil {
			return hotel, err
		}

		featureHotel.Qty = j
		featureHotel.ImageUrl = input.Features[i][2]
		featureHotel.HotelID = id
		_, err = s.FeatureRepository.SaveHotel(featureHotel)
		if err != nil{
			return hotel, err
		}
	}

	
	_, err = s.HotelRepository.Update(hotel)
	if err != nil {
		fmt.Println("Errornya disini")
		return hotel, err
	}
	hotel, err = s.HotelRepository.FindByIDDetail(id)
	if err != nil {
		return hotel, err
	}
	return hotel, nil

}
// func (s *HotelService) GetHotel(limit, page int)([]domain.Hotel, error){
// 	var Hotel []domain.Hotel
// 	Hotel, err := s.HotelRepository.GetHotel(limit, page)
// 	if err != nil {
// 		fmt.Println("Errornya disini")
// 		return Hotel, err
// 	}
// 	return Hotel, nil
// }