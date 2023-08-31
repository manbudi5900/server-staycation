package domain

import "time"

type Category struct {
	ID int `gorm:"autoIncrement:1;primary_key"`
	Name string
	Products []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}