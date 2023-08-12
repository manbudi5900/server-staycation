package formatter

import (
	"staycation/domain"
)

type ProductImageFormatter struct {
	ID       int `json:"id"`
	Url     string `json:"url"`
	
}

func FormatProductImage(productImage domain.ProductImage) ProductImageFormatter {
	formatter := ProductImageFormatter{
		ID:       productImage.ID,
		Url:     productImage.Url,
		
	}

	return formatter
}