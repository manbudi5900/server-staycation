package domain

import "time"

type Category struct {
	ID int `gorm:"autoIncrement:1;primary_key"`
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}