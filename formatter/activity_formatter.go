package formatter

import "staycation/domain"

type ActivityFormatter struct {
	ID       int `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"qty"`
	ImageUrl     string `json:"image_url"`
	
}

func FormatActivity(activity domain.Activity) ActivityFormatter {
	formatter := ActivityFormatter{
		ID:       activity.ID,
		Name:     activity.Name,
		Type:     activity.Type,
		ImageUrl:     activity.ImageUrl,
		
	}

	return formatter
}