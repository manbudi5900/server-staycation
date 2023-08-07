package routes

import (
	"staycation/handler"
	"staycation/middleware"
	"staycation/repository"
	"staycation/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ActivityRoute(routes (*gin.Engine) ,api handler.ActivityHandler, dbConfig *gorm.DB) {
	// routes := gin.Default()
	userRepository := repository.NewUserRepository(dbConfig)

	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService()
	prd := routes.Group("/api/v1")
	{
		prd.POST("/activity/store", middleware.AuthAdminMiddleware(authService, userService) ,api.SaveActivity)
		prd.GET("/activity", middleware.AuthAdminMiddleware(authService, userService) ,api.GetActivity)
		prd.PUT("/activity/:activityId",middleware.AuthAdminMiddleware(authService, userService) ,api.UpdateActivity)
	}
}