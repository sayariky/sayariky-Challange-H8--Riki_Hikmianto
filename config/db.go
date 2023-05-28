package config

import (
	"go-trial-class/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	conn := "host=localhost port=5432 user=postgres password=060818 dbname=contoh_db sslmode=disable"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	} else {
		log.Println("Database connected")
		DB = db
	}
	db.AutoMigrate(&entity.Product{}, &entity.Order{}, &entity.User{})
}
