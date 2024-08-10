package usecase

import (
	"MatchingApp/internal/model"
	"github.com/gofrs/uuid/v5"
)

type UserUseCase interface {
	CreateUser(user *model.User) error
	DeleteUserByID(id uuid.UUID) error
	FindUserByID(id uuid.UUID) (*model.User, error)
	FindUserByUsername(username string) (*model.User, error)
	FindMatch(id uuid.UUID) *model.User
}
