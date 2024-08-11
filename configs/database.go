package configs

import (
	"golang_jwt_copy/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := "host=localhost user=postgres password=Whobay123@ dbname=go_jwt port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed Connecting to Database")
	}

	db.AutoMigrate(&models.User{})

	DB = db
	log.Println("Succesfully Connected to Database")
}
