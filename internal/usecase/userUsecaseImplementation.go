package usecase

import (
	"MatchingApp/internal/model"
	"MatchingApp/internal/repository"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"log"
)

type UserUseCaseImplementation struct {
	repository repository.UserRepository
}

func NewUserUseCaseImplementation(repository repository.UserRepository) *UserUseCaseImplementation {
	return &UserUseCaseImplementation{repository: repository}
}

func (uc *UserUseCaseImplementation) CreateUser(user *model.User) error {
	if user.Email == "" || user.Password == nil || user.Username == "" {
		return fmt.Errorf("fileds can't be empty")
	}

	if foundedUser, err := uc.FindUserByUsername(user.Username); foundedUser != nil || err != nil {
		return fmt.Errorf("user with username %s already exists, try another username", user.Username)
	}

	err := uc.repository.CreateUser(user)
	return err
}

func (uc *UserUseCaseImplementation) DeleteUserByID(id uuid.UUID) error {
	foundedUser, err := uc.FindUserByID(id)
	if err != nil {
		return err
	} else if foundedUser == nil {
		log.Printf("can't delete user with id %s, there is no such user", id)
		return fmt.Errorf("can't delete user with id %s, there is no such user", id)
	}
	err = uc.repository.DeleteUserByID(id)
	return err
}

func (uc *UserUseCaseImplementation) FindUserByID(id uuid.UUID) (*model.User, error) {
	foundedUser, err := uc.repository.FindUserByID(id)
	if err != nil {
		return nil, err
	}
	return foundedUser, nil
}

func (uc *UserUseCaseImplementation) FindUserByUsername(username string) (*model.User, error) {
	foundedUser, err := uc.repository.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return foundedUser, nil
}

func (uc *UserUseCaseImplementation) FindMatch(id uuid.UUID) *model.User {
	return uc.repository.FindMatch(id)
}
