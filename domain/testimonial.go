package domain

import "time"

type Testimonial struct {
	ID int `gorm:"primary_key;"`
	Name string
	ImageUrl string
	Rate float32
	Content string
	FamilyName string
	FamilyOccupation string
	CreatedAt time.Time
	UpdatedAt time.Time
}
