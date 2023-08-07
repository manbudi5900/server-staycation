package routes

import (
	db2 "staycation/db"
	"staycation/injection"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine{
	dbConfig := db2.InitDB()

	productAPI := injection.InitProductAPI(dbConfig)
	categoryAPI := injection.InitCategoryAPI(dbConfig)
	userAPI := injection.InitUserAPI(dbConfig)
	featureAPI := injection.InitFeatureAPI(dbConfig)
	activityAPI := injection.InitActivityAPI(dbConfig)

	routes := gin.Default()

	// set logger
	routes.Use(gin.Logger())

	// Gzip Compression
	routes.Use(gin.Recovery())

	ProductRoute(routes, productAPI, dbConfig)
	CategoryRoute(routes, categoryAPI, dbConfig)
	UserRoute(routes, userAPI, dbConfig)
	FeatureRoute(routes, featureAPI, dbConfig)
	ActivityRoute(routes, activityAPI, dbConfig)


	return routes
}