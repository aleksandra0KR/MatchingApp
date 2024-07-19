package model

import (
	"github.com/gofrs/uuid/v5"
)

type Playlist struct {
	ID               uuid.UUID `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v4()"`
	PlaylistKey      string    `json:"playlist_key,omitempty" gorm:"type:text"`
	UserID           uuid.UUID `json:"userId,omitempty" gorm:"type:uuid;references:ID;foreignKey:ID"`
	UserName         string    `json:"userName,omitempty" gorm:"type:varchar(255)"`
	Danceability     float64   `json:"danceability,omitempty" gorm:"type:DOUBLE PRECISION;"`
	Energy           float64   `json:"energy,omitempty" gorm:"type:DOUBLE PRECISION;"`
	Key              float64   `json:"key,omitempty" gorm:"type:DOUBLE PRECISION;"`
	Loudness         float64   `json:"loudness,omitempty" gorm:"type:DOUBLE PRECISION;"`
	Mode             float64   `json:"mode,omitempty" gorm:"type:DOUBLE PRECISION;"`
	Speechiness      float64   `json:"speechiness,omitempty" gorm:"type:DOUBLE PRECISION;"`
	Acousticness     float64   `json:"acousticness,omitempty" gorm:"type:DOUBLE PRECISION;"`
	Instrumentalness float64   `json:"instrumentalness,omitempty" gorm:"type:DOUBLE PRECISION;"`
	Liveness         float64   `json:"liveness,omitempty" gorm:"type:DOUBLE PRECISION;"`
	Valence          float64   `json:"valence,omitempty" gorm:"type:DOUBLE PRECISION;"`
	Tempo            float64   `json:"tempo,omitempty" gorm:"type:DOUBLE PRECISION;"`
	Duration_ms      float64   `json:"duration_ms,omitempty" gorm:"type:DOUBLE PRECISION;"`
	Time_signature   float64   `json:"time_signature,omitempty" gorm:"type:DOUBLE PRECISION;"`
}
