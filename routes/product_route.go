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
		// prd.GET("/",func(c *gin.Context) {
		// 	c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
		//   })
		prd.POST("/upload-image/:folder", middleware.AuthAdminMiddleware(authService, userService), api.UploadImage)
		prd.POST("/product/store", middleware.AuthAdminMiddleware(authService, userService) ,api.Save)
		prd.GET("/landing-page", api.GetLandingPage, middleware.CORSMiddleware())
		prd.GET("/product-detail/:slug", api.GetProductDetail)

		
	}
}