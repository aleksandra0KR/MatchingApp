package repository

import (
	"MatchingApp/internal/model"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
	"log"
)

type PlaylistPostgresRepository struct {
	db *gorm.DB
}

func NewPlaylistPostgresRepository(db *gorm.DB) *PlaylistPostgresRepository {
	return &PlaylistPostgresRepository{db: db}
}

func (r *PlaylistPostgresRepository) CreatePlayList(playlist *model.PlayList) {
	result := r.db.Create(&playlist)

	if result.Error != nil {
		log.Printf("Failed to create a new playlist: %v", result.Error)
	}
}

func (r *PlaylistPostgresRepository) DeletePlayListByID(id uuid.UUID) {
	result := r.db.Delete(id)

	if result.Error != nil {
		log.Printf("Failed to delete a playlist with id %d: %v", id, result.Error)
	}
}

func (r *PlaylistPostgresRepository) FindPlayListByID(id uuid.UUID) *model.PlayList {
	playlist := &model.PlayList{PlaylistId: id}
	result := r.db.First(playlist)
	if result.Error != nil {
		log.Printf("Failed to find a user with id %s: %v", id, result.Error)
		return nil
	}
	return playlist
}
