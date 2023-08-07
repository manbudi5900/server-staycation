package domain

import "time"

type Activity struct {
	ID int `gorm:"autoIncrement:1;primary_key"`
	Name string
	Type string
	ImageUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
}