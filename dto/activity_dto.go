package dto

type ActivityInput struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
	ImageUrl string `json:"image_url" binding:"required"`

}
