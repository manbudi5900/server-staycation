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
	CreatedAt time.Time
	UpdatedAt time.Time
}