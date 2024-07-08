package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository
	PlaylistRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:     NewUserPostgresRepository(db),
		PlaylistRepository: NewPlaylistPostgresRepository(db),
	}
}
