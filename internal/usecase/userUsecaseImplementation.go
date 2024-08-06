package usecase

import (
	"MatchingApp/internal/model"
	"MatchingApp/internal/repository"
	"fmt"
	"github.com/gofrs/uuid/v5"
)

type UserUseCaseImplementation struct {
	repository repository.UserRepository
}

func NewUserUseCaseImplementation(repository repository.UserRepository) *UserUseCaseImplementation {
	return &UserUseCaseImplementation{repository: repository}
}

func (uc *UserUseCaseImplementation) CreateUser(user *model.User) {
	if user.Email == "" || user.Password == nil || user.Username == "" {
		fmt.Errorf("all fileds can't be empty")
	}
	uc.repository.CreateUser(user)
}

func (uc *UserUseCaseImplementation) DeleteUserByID(id uuid.UUID) {
	uc.repository.DeleteUserByID(id)
}
func (uc *UserUseCaseImplementation) FindUserByID(id uuid.UUID) *model.User {
	foundedUser := uc.repository.FindUserByID(id)

	return foundedUser
}

func (uc *UserUseCaseImplementation) FindUserByUsername(username string) *model.User {
	foundedUser := uc.repository.FindUserByUsername(username)

	return foundedUser
}
func (uc *UserUseCaseImplementation) FindMatch(id uuid.UUID) *model.User {
	return uc.repository.FindMatch(id)
}
