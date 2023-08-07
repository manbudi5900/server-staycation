package user

type RegisterUserInput struct {
	Name string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"phone" binding:"required"`
	Province string `json:"province" binding:"required"`
	City string `json:"city" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}