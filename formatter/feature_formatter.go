package formatter

import "staycation/domain"

type FeatureFormatter struct {
	ID       int `json:"id"`
	Name     string `json:"name"`
	Qty     int `json:"qty"`
	ImageUrl     string `json:"image_url"`
	
}

func FormatFeature(feature domain.Feature) FeatureFormatter {
	formatter := FeatureFormatter{
		ID:       feature.ID,
		Name:     feature.Name,
		Qty:     feature.Qty,
		ImageUrl:     feature.ImageUrl,
		
	}

	return formatter
}
