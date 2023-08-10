package config

import (
	"log"
	"os"
	"staycation/domain"

	"github.com/joho/godotenv"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB  {
	processENV()

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")


	dsn := "host="+dbHost+" user="+dbUsername+" password="+dbPassword+" dbname="+dbName+" port="+dbPort+" sslmode=disable TimeZone=Asia/Jakarta"
	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true,})
	if err != nil {
		log.Fatal(err.Error())
	}

	migrateDDL(db)

	return db
}

func migrateDDL(db *gorm.DB) {
	db.AutoMigrate(&domain.ProductImage{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Product{})
}

func processENV() {

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error("Error loading env file")
	}
}