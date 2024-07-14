package repository

import (
	"MatchingApp/internal/model"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
	"log"
)

type UserPostgresRepository struct {
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) *UserPostgresRepository {
	return &UserPostgresRepository{db: db}
}

func (r *UserPostgresRepository) CreateUser(user *model.User) {
	result := r.db.Create(&user)

	if result.Error != nil {
		log.Printf("Failed to create a new user: %v", result.Error)
	}
}

func (r *UserPostgresRepository) DeleteUserByID(id uuid.UUID) {
	result := r.db.Delete(id)

	if result.Error != nil {
		log.Printf("Failed to delete a user with id %d: %v", id, result.Error)
	}
}

func (r *UserPostgresRepository) FindUserByID(id uuid.UUID) *model.User {
	var user model.User
	result := r.db.First(&user, id)
	fmt.Println(user)
	if result.Error != nil {
		log.Printf("Failed to find a user with id %s: %v", id, result.Error)
		return nil
	}
	return &user
}
func (r *UserPostgresRepository) FindUserByUsername(username string) *model.User {
	var user model.User
	result := r.db.First(&user, "username = ?", username)
	if result.Error != nil {
		log.Printf("Failed to find a user with usernae,e %s: %v", user, result.Error)
		return nil
	}
	return &user
}

// FindMatch TODO
func (r *UserPostgresRepository) FindMatch(id uuid.UUID) *model.User {
	return nil
}
