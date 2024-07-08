package repository

import (
	"MatchingApp/internal/model"
	"github.com/gofrs/uuid/v5"
)

type UserRepository interface {
	CreateUser(user *model.User)
	DeleteUserByID(id uuid.UUID)
	FindUserByID(id uuid.UUID) *model.User
	FindMatch(id uuid.UUID) *model.User
}
