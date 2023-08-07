package dto

type FeatureInput struct {
	Name string `json:"name" binding:"required"`
	Qty int `json:"qty" binding:"required"`
	ImageUrl string `json:"image_url" binding:"required"`

}
