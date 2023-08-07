package routes

import (
	"staycation/handler"
	"staycation/middleware"
	"staycation/repository"
	"staycation/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRoute(routes (*gin.Engine) ,api handler.ProductHandler, dbConfig *gorm.DB) {
	// routes := gin.Default()
	userRepository := repository.NewUserRepository(dbConfig)

	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService()
	prd := routes.Group("/api/v1")
	{
		prd.POST("/upload-image/:folder", middleware.AuthAdminMiddleware(authService, userService), api.UploadImage)
		prd.POST("/product/store", middleware.AuthAdminMiddleware(authService, userService) ,api.Save)
		
	}
}