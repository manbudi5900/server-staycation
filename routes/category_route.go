package routes

import (
	"staycation/handler"
	"staycation/middleware"
	"staycation/repository"
	"staycation/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CategoryRoute(routes (*gin.Engine) ,api handler.CategoryProductHandler, dbConfig *gorm.DB) {
	// routes := gin.Default()
	userRepository := repository.NewUserRepository(dbConfig)

	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService()
	
	
	prd := routes.Group("/api/v1")
	{
		prd.POST("/category/store", middleware.AuthAdminMiddleware(authService, userService),api.SaveCategory)
		prd.GET("/category", middleware.AuthAdminMiddleware(authService, userService) ,api.GetCategory)
		prd.PUT("/category/:categoryId",middleware.AuthAdminMiddleware(authService, userService) ,api.UpdateCategory)
		
	}
}