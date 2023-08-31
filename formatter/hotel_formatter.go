package formatter

import (
	"staycation/domain"
)

type HotelDetailFormatter struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Slug string `json:"slug"`
	Country string `json:"country"`
	Province string `json:"province"`
	City string `json:"city"`
	Logo string `json:"logo"`
	Desc string `json:"description"`
	ProductImage []ProductImageFormatter `json:"images"`
	Feature []FeatureFormatter `json:"features"`
}
type HotelFormatter struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Slug string `json:"slug"`
	Country string `json:"country"`
	Province string `json:"province"`
	City string `json:"city"`
	Logo string `json:"logo"`
	Desc string `json:"description"`
}

func FormatHotelDetail(hotel domain.Hotel) HotelDetailFormatter {
	var formatterItemImage []ProductImageFormatter


	for _,items := range hotel.ProductImageHotel {
		formatterProductImage :=  FormatProductImage(items)
		formatterItemImage = append(formatterItemImage, formatterProductImage)

	}
	if formatterItemImage == nil {
		formatterItemImage = []ProductImageFormatter{}
	}

	var formatterItemFeature []FeatureFormatter


	for _,items := range hotel.Feature {
		formatterHotelFeature :=  FormatFeature(items)
		formatterItemFeature = append(formatterItemFeature, formatterHotelFeature)

	}
	if formatterItemFeature == nil {
		formatterItemFeature = []FeatureFormatter{}
	}
	formatter := HotelDetailFormatter{
		ID:       hotel.ID,
		Name:     hotel.Name,
		Slug:    hotel.Slug,
		Province:    hotel.Province,
		City: hotel.City,
		Logo:    hotel.Logo,
		Desc:    hotel.Desc,
		ProductImage: formatterItemImage,
		Feature: formatterItemFeature,
	}

	return formatter
}

func FormatHotel(hotel []domain.Hotel) []HotelFormatter {
	var fortmatters = []HotelFormatter{}

	for _,items := range hotel {
		formatter := HotelFormatter{
			ID:       items.ID,
			Name:     items.Name,
			Slug:    items.Slug,
			Province:    items.Province,
			City: items.City,
			Logo:    items.Logo,
			Desc:    items.Desc,
		}
		fortmatters = append(fortmatters, formatter)
	}

	

	return fortmatters
}