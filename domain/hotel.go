package domain

import "time"

type Hotel struct {
	ID string `gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Logo string `json:"logo"`
	Province string `json:"province"`
	City string `json:"city"`
	Desc string `json:"desc"`
	ProductImageHotel []ProductImage `json:"product_images"`
	Feature []Feature `json:"features"`
	CreatedAt time.Time
	UpdatedAt time.Time
}