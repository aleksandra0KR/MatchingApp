package repository

import (
	"MatchingApp/internal/model"
	"github.com/gofrs/uuid/v5"
)

type PlaylistRepository interface {
	CreatePlaylist(playlist *model.Playlist)
	DeletePlaylistByID(id uuid.UUID)
	FindPlaylistByID(id uuid.UUID) *model.Playlist
}
