package repository

import (
	"MatchingApp/internal/model"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
	"log"
)

type UserPostgresRepository struct {
	db *gorm.DB
}

func NewQuestPostgresRepository(db *gorm.DB) *UserPostgresRepository {
	return &UserPostgresRepository{db: db}
}

func (r *UserPostgresRepository) CreateUser(user *model.User) {
	result := r.db.Create(&user)

	if result.Error != nil {
		log.Printf("Failed to create a new user: %v", result.Error)
	}
}

func (r *UserPostgresRepository) DeleteUser(id uuid.UUID) {
	result := r.db.Delete(id)

	if result.Error != nil {
		log.Printf("Failed to delete a user with id %d: %v", id, result.Error)
	}
}

func (r *UserPostgresRepository) FindUserByID(id uuid.UUID) *model.User {
	user := &model.User{UserId: id}
	result := r.db.First(user)
	if result.Error != nil {
		log.Printf("Failed to find a user with id %s: %v", id, result.Error)
		return nil

	}
	return user
}

// FindMatch TODO
func (r *UserPostgresRepository) FindMatch(id uuid.UUID) *model.User {
	return nil
}
