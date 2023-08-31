package domain

import (
	"time"
)

type Transaction struct {
	ID string `gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	UserID string
	ProductID string
	Status int
	Amount int
	Code string
	PaymentURL string
	User User
	Product Product
	CreatedAt time.Time
	UpdatedAt time.Time
}
type MidtransConfig struct {
	ClientKey string
	ServerKey string
	APIEnv    string
}
