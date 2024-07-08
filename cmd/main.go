package main

import (
	"MatchingApp/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "host=localhost user=postgres dbname=MatchingApp password=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("cant connect to postgresql: %v", err)
		return
	}

	err = db.AutoMigrate(&model.User{}, &model.PlayList{})
	if err != nil {
		log.Fatalf("cant initialize to postgresql: %v", err)
		return
	}
}
