package dto

type HotelInput struct {
	Name string `json:"name" binding:"required"`
	Logo string `json:"logo" binding:"required"`
	Desc string `json:"desc" binding:"required"`
	Province string `json:"province" binding:"required"`
	City string `json:"city" binding:"required"`
	Images []string `json:"images"`
	Features [][]string `json:"features"`
}
