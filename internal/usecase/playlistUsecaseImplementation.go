package usecase

import (
	"MatchingApp/internal/model"
	"MatchingApp/internal/repository"
	"github.com/gofrs/uuid/v5"
)

type PlaylistUseCaseImplementation struct {
	repository repository.PlaylistRepository
}

func NewPlaylistUseCaseImplementation(repository repository.PlaylistRepository) *PlaylistUseCaseImplementation {
	return &PlaylistUseCaseImplementation{repository: repository}
}

func (plu *PlaylistUseCaseImplementation) CreatePlaylist(playlist *model.Playlist) {
	plu.repository.CreatePlaylist(playlist)
}
func (plu *PlaylistUseCaseImplementation) DeletePlaylistByID(id uuid.UUID) {
	plu.repository.DeletePlaylistByID(id)
}
func (plu *PlaylistUseCaseImplementation) FindPlaylistByID(id uuid.UUID) *model.Playlist {
	return plu.repository.FindPlaylistByID(id)
}
