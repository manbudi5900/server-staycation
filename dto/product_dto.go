package dto

type ProductInput struct {
	Name string `json:"name" binding:"required"`
	CategoryID string `json:"category_id" binding:"required"`
	Price int `json:"price" binding:"required"`
	Country string `json:"country" binding:"required"`
	Description string `json:"description" binding:"required"`
	Province string `json:"province" binding:"required"`
	City string `json:"city" binding:"required"`
	Url []string `json:"url" binding:"required"`
}
