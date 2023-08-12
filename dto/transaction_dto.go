package dto

import "staycation/domain"

type TransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	ProductID string `json:"product_id" binding:"required"`
	User       domain.User
}
