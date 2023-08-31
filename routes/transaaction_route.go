package routes

import (
	"staycation/handler"
	"staycation/middleware"
	"staycation/repository"
	"staycation/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TransactionRoute(routes (*gin.Engine) ,api handler.TransactionHandler, dbConfig *gorm.DB) {
	userRepository := repository.NewUserRepository(dbConfig)

	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService()
	prd := routes.Group("/api/v1")
	{
		prd.POST("/transaction/store", middleware.AuthCustomerMiddleware(authService, userService),middleware.DBTransactionMiddleware(dbConfig), api.CreateTransaction)
		// prd.POST("/product/store", middleware.AuthCustomerMiddleware(authService, userService) ,api.Save)
		prd.GET("/transactions", middleware.AuthCustomerMiddleware(authService, userService), middleware.DBTransactionMiddleware(dbConfig), api.GetUserTransactions)
		prd.POST("/transactions/notification", api.GetNotification)


		
	}
}