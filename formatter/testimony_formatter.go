package formatter

import (
	"staycation/domain"
)

type TestimonialFormatter struct {
	ID       int `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	Rate float32 `json:"rate"`
	Content string `json:"content"`
	FamilyName string `json:"familyName"`
	FamilyOccupation string `json:"familyOccupation"`
	
}

func FormatTestimonial(Testimonial domain.Testimonial) TestimonialFormatter {
	formatter := TestimonialFormatter{
		ID:       Testimonial.ID,
		Name:     Testimonial.Name,
		ImageUrl:     Testimonial.ImageUrl,
		Rate:     Testimonial.Rate,
		Content:     Testimonial.Content,
		FamilyName:     Testimonial.FamilyName,
		FamilyOccupation:     Testimonial.FamilyOccupation,

		
	}

	return formatter
}