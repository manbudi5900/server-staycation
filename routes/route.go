package routes

import (
	db2 "staycation/config"
	"staycation/injection"
	"fmt"
	"net/http"


	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"

)

func Init() *gin.Engine{
	dbConfig := db2.InitDB()

	productAPI := injection.InitProductAPI(dbConfig)
	categoryAPI := injection.InitCategoryAPI(dbConfig)
	userAPI := injection.InitUserAPI(dbConfig)
	featureAPI := injection.InitFeatureAPI(dbConfig)
	activityAPI := injection.InitActivityAPI(dbConfig)

	transactionAPI := injection.InitTransactionAPI(dbConfig)

	routes := gin.Default()

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*"} // Replace with your frontend's URL
	// config.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}
	// config.AllowHeaders = []string{"Content-Type"}
	// config.AllowHeaders = []string{"Content-Type", "Access-Control-Allow-Origin"} // Add "Access-Control-Allow-Origin"
	// config.AllowHeaders = []string{"Content-Type", "Access-Control-Allow-Headers"} // Add "Access-Control-Allow-Origin"


	// routes.Use(cors.New(config))
    // routes.Use(cors.Default())


	// set logger
	routes.Use(gin.Logger())

	

	// Gzip Compression
	routes.Use(gin.Recovery())
	routes.Static("/images", "./images")
	routes.GET("/",func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
	  })

	fmt.Println("route")

	ProductRoute(routes, productAPI, dbConfig)
	CategoryRoute(routes, categoryAPI, dbConfig)
	FeatureRoute(routes, featureAPI, dbConfig)
	ActivityRoute(routes, activityAPI, dbConfig)
	TransactionRoute(routes, transactionAPI, dbConfig)
	UserRoute(routes, userAPI, dbConfig)


	return routes
}