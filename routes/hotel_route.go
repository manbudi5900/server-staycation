package routes

import (
	"staycation/handler"
	"staycation/middleware"
	"staycation/repository"
	"staycation/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HotelRoute(routes (*gin.Engine) ,api handler.HotelHandler, dbConfig *gorm.DB) {

	userRepository := repository.NewUserRepository(dbConfig)

	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService()
	prd := routes.Group("/api/v1")
	{
		prd.POST("/hotel/store", middleware.AuthAdminMiddleware(authService, userService) ,api.Save)
		prd.GET("/hotel", middleware.AuthAdminMiddleware(authService, userService) ,api.GetHotel)
		prd.PUT("/hotel/:hotelId",middleware.AuthAdminMiddleware(authService, userService) ,api.Update)
	}
}