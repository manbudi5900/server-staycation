package injection

import (
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