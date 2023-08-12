package config

import (
	"log"
	"staycation/domain"

	"github.com/joho/godotenv"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB  {
	processENV()

	// dbUsername := os.Getenv("DB_USERNAME")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_DATABASE")

	var dbUsername = "edwzgbymhxftlx"
	var dbPassword = "2ff13d0e42b4c64cebc75d88f9e6ada58b1b63b6470654ceb47d4fdfc1556565"
	var dbHost = "ec2-52-215-68-14.eu-west-1.compute.amazonaws.com"
	var dbPort = "5432"
	var dbName ="ddu8acp2fpitg6"


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