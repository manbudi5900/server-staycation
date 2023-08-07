package user

import (
	"time"
)

type User struct {
	ID string `gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	Role string
	Name string
	Status int
	Username string
	Email string
	Phone string
	Province string
	City string
	Avatar string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}