package domain

import "time"

type Product struct {
	ID string `gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	Name string
	Slug string
	CategoryID string
	Price int
	Country string
	Province string
	City string
	Description string
	IsPopuler bool
	Unit string `json:"unit"`
	Type string `json:"type"`
	IsBooking int `json:"is_boking"`
	ProductImage []ProductImage `json:"product_images"`
	Activity []Activity 
	Feature []Feature
	Category Category
	CreatedAt time.Time
	UpdatedAt time.Time
}