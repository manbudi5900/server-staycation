package dto

type ProductInput struct {
	Name string `json:"name" binding:"required"`
	CategoryID int `json:"category_id" binding:"required"`
	Price int `json:"price" binding:"required"`
	Country string `json:"country" binding:"required"`
	Description string `json:"description" binding:"required"`
	Province string `json:"province" binding:"required"`
	City string `json:"city" binding:"required"`
	IsBooking int `json:"is_boking"`
	Url []string `json:"url" binding:"required"`
}
