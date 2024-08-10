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

func NewUserPostgresRepository(db *gorm.DB) *UserPostgresRepository {
	return &UserPostgresRepository{db: db}
}

func (r *UserPostgresRepository) CreateUser(user *model.User) error {
	result := r.db.Create(&user)

	if result.Error != nil {
		log.Printf("Failed to insert new user to db: %v", result.Error)
		return result.Error
	}
	return nil
}

func (r *UserPostgresRepository) DeleteUserByID(id uuid.UUID) error {
	result := r.db.Delete(id)

	if result.Error != nil {
		log.Printf("Failed to delete a user with id from db%d: %v", id, result.Error)
		return result.Error
	}
	return nil
}

func (r *UserPostgresRepository) FindUserByID(id uuid.UUID) (*model.User, error) {
	var user model.User
	result := r.db.First(&user, id)

	if result.Error != nil {
		log.Printf("Failed to find a user with id %s in db: %v", id, result.Error)
		return nil, result.Error
	}
	return &user, nil
}
func (r *UserPostgresRepository) FindUserByUsername(username string) (*model.User, error) {
	var user model.User
	result := r.db.First(&user, "username = ?", username)

	if result.Error != nil {
		log.Printf("Failed to find a user with username %s in db: %v", user, result.Error)
		return nil, result.Error
	}
	return &user, nil
}

// FindMatch TODO
func (r *UserPostgresRepository) FindMatch(id uuid.UUID) *model.User {
	return nil
}
