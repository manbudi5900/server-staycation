package formatter

import (
	"staycation/domain"
)

type ProductFormatter struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Slug string `json:"slug"`
	CategoryID string `json:"category_id"`
	Price int `json:"price"`
	Country string `json:"country"`
	Province string `json:"province"`
	City string `json:"city"`
	ImageURL string `json:"image_url"`
	Description string `json:"description"`
}

func FormatProduct(product domain.Product) ProductFormatter {
	formatter := ProductFormatter{
		ID:       product.ID,
		Name:     product.Name,
		Slug:    product.Slug,
		CategoryID:    product.CategoryID,
		Price:    product.Price,
		Country: product.Country,
		Province:    product.Province,
		City: product.City,
		Description:    product.Description,
	}

	return formatter
}