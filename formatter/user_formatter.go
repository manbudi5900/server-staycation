package formatter

import (
	"staycation/domain"
)

type UserFormatter struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Token    string `json:"token"`
	ImageURL string `json:"image_url"`
}

func FormatUser(user domain.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Token:    token,
		ImageURL: user.Avatar,
	}

	return formatter
}