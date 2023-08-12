package injection

import (
	"log"
	"staycation/config"
	"staycation/handler"
	"staycation/repository"
	"staycation/service"

	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitProductAPI(db *gorm.DB) handler.ProductHandler {
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productAPI := handler.NewProductHandler(productService, service.NewAuthService())
	return productAPI
}
func InitCategoryAPI(db *gorm.DB) handler.CategoryProductHandler {
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryAPI := handler.NewCategoryProductHandler(categoryService, service.NewAuthService())
	return categoryAPI
}
func InitUserAPI(db *gorm.DB) handler.UserHandler {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userAPI := handler.NewUserHandler(userService, service.NewAuthService())
	return userAPI
}
func InitFeatureAPI(db *gorm.DB) handler.FeatureHandler {
	featureRepository := repository.NewFeatureRepository(db)
	featureService := service.NewFeatureService(featureRepository)
	featureAPI := handler.NewFeatureHandler(featureService, service.NewAuthService())
	return featureAPI
}
func InitActivityAPI(db *gorm.DB) handler.ActivityHandler {
	activityRepository := repository.NewActivityRepository(db)
	activityService := service.NewActivityService(activityRepository)
	activityAPI := handler.NewActivityHandler(activityService, service.NewAuthService())
	return activityAPI
}
func InitTransactionAPI(db *gorm.DB) handler.TransactionHandler {
	config1, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err.Error())
	}
	productRepository := repository.NewProductRepository(db)
	paymentRepository := repository.NewPaymentRepository(config1.Midtrans)
	paymentService := service.NewPaymentService(paymentRepository)

	activityRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(activityRepository,paymentService ,productRepository)
	activityAPI := handler.NewTransactionHandler(transactionService, service.NewAuthService())
	return activityAPI
}