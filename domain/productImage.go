package domain

import "time"

type ProductImage struct {
	ID int `gorm:"autoIncrement:1;primary_key"`
	ProductID string
	Url string
	CreatedAt time.Time
	UpdatedAt time.Time
}