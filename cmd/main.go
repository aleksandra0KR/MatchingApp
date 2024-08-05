package main

import (
	handler2 "MatchingApp/internal/handler"
	"MatchingApp/internal/model"
	"MatchingApp/internal/repository"
	"MatchingApp/internal/server"
	"MatchingApp/internal/usecase"
	"MatchingApp/kafka"
	"github.com/gofrs/uuid/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"html/template"
	"log"
)

var tpl *template.Template

func main() {

	tpl, _ = template.ParseGlob("templates/*.html")

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

	useCase := usecase.NewUseCase(repo)
	producer := kafka.SetUpProducer()
	consumer := kafka.SetUpConsumer()
	handler := handler2.NewHandler(useCase, tpl, consumer, producer)

	var srv server.Server
	err = srv.Run("8080", handler.Handle())
}

/*
test data

username: eric
playlist id : 2DLfBC25Ven3wR1Sujbd8R
*/
