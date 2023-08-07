package dto

type CategoryInput struct {
	Name string `json:"name" binding:"required"`
}
