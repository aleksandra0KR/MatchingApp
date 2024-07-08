package main

import (
	"MatchingApp/internal/model"
	"MatchingApp/internal/repository"
	"github.com/gofrs/uuid/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "host=localhost port=5436 user=postgres dbname=postgres password=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("cant connect to postgresql: %v", err)
		return
	}

	err = db.AutoMigrate(&model.User{}, &model.Playlist{})
	if err != nil {
		log.Fatalf("cant initialize to postgresql: %v", err)
		return
	}
	repo := repository.NewRepository(db)
	repo.UserRepository.CreateUser(&model.User{
		Email:    "ewd",
		Username: "wderf",
		Match:    make([]model.User, 0),
	})

	uu, _ := uuid.FromString("3e7957b2-a10c-469c-8473-cd3644a6f7cf")
	repo.PlaylistRepository.CreatePlaylist(&model.Playlist{UserID: uu, Energy: 1})
}
