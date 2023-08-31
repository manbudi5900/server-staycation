package domain

import "time"

type Feature struct {
	ID int `gorm:"autoIncrement:1;primary_key"`
	Name string
	Qty int
	ImageUrl string
	ProductID string
	HotelID string
	CreatedAt time.Time
	UpdatedAt time.Time
}