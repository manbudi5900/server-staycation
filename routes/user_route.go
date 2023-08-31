package routes

import (
	"staycation/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoute(routes (*gin.Engine) ,api handler.UserHandler, dbConfig *gorm.DB) {
	// routes := gin.Default()
	// userRepository := repository.NewUserRepository(dbConfig)

	// userService := service.NewUserService(userRepository)
	// authService := service.NewAuthService()
	prd := routes.Group("/api/v1")
	{
		prd.POST("/register", api.RegisterUser)
		prd.POST("/login", api.LoginUser)
		prd.POST("cek-token", api.CekToken)
	}
}