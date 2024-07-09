package usecase

import (
	"MatchingApp/internal/model"
	"github.com/gofrs/uuid/v5"
)

type PlaylistUseCase interface {
	CreatePlaylist(playlist *model.Playlist)
	DeletePlaylistByID(id uuid.UUID)
	FindPlaylistByID(id uuid.UUID) *model.Playlist
}
