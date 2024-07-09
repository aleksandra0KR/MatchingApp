package usecase

import "MatchingApp/internal/repository"

type UseCase struct {
	UserUseCase
	PlaylistUseCase
}

func NewUseCase(repository *repository.Repository) *UseCase {
	return &UseCase{
		UserUseCase:     NewUserUseCaseImplementation(repository.UserRepository),
		PlaylistUseCase: NewPlaylistUseCaseImplementation(repository.PlaylistRepository),
	}
}
