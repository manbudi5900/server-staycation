package routes

import (
	"staycation/handler"
	"staycation/middleware"
	"staycation/repository"
	"staycation/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FeatureRoute(routes (*gin.Engine) ,api handler.FeatureHandler, dbConfig *gorm.DB) {
	// routes := gin.Default()
	userRepository := repository.NewUserRepository(dbConfig)

	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService()
	prd := routes.Group("/api/v1")
	{
		prd.POST("/feature/store", middleware.AuthAdminMiddleware(authService, userService) ,api.SaveFeature)
		prd.GET("/feature", middleware.AuthAdminMiddleware(authService, userService) ,api.GetFeature)
		prd.PUT("/feature/:featureId",middleware.AuthAdminMiddleware(authService, userService) ,api.UpdateFeature)
	}
}