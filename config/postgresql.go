package config

import (
	"fmt"
	"log"
	"myproject/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBP *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=1310 dbname=todoapp sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connect database: ", err)
	}

	db.AutoMigrate(&models.User{}, &models.Task{})
	DBP = db
	fmt.Println("Connect postgres!")

	return db
}
