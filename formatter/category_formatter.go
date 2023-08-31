package formatter

import "staycation/domain"

type CategoryFormatter struct {
	ID       int `json:"id"`
	Name     string `json:"name"`
	
}

func FormatCategroy(category domain.Category) CategoryFormatter {
	formatter := CategoryFormatter{
		ID:       category.ID,
		Name:     category.Name,
		
	}

	return formatter
}
