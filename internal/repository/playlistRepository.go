package repository

import (
	"MatchingApp/internal/model"
	"github.com/gofrs/uuid/v5"
)

type PlaylistRepository interface {
	CreatePlaylist(playlist *model.PlayList)
	DeletePlaylistByID(id uuid.UUID)
	FindPlaylistByID(id uuid.UUID) *model.PlayList
}
