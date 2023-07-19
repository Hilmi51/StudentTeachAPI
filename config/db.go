package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"student-teach-api/models"
	"time"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("host=postgres user=postgres password=postgres dbname=student_api port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	start := time.Now()
	for sqlDB.Ping() != nil {
		if start.After(start.Add(10 * time.Second)) {
			fmt.Println("Failed to connect DB after 10 seconds")
			break
		}
	}

	fmt.Println("Connected to DB: ", sqlDB.Ping() == nil)
	db.AutoMigrate(&models.User{})
	DB = db
}
