package model

import (
	"github.com/gofrs/uuid/v5"
)

type Playlist struct {
	ID               uuid.UUID `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID           uuid.UUID `json:"userId,omitempty" gorm:"type:uuid;references:ID;foreignKey:ID"`
	Danceability     float64   `json:"danceability,omitempty" gorm:"type:decimal(7,6);"`
	Energy           float64   `json:"energy,omitempty" gorm:"type:decimal(7,6);"`
	Key              float64   `json:"key,omitempty" gorm:"type:decimal(7,6);"`
	Loudness         float64   `json:"loudness,omitempty" gorm:"type:decimal(7,6);"`
	Mode             float64   `json:"mode,omitempty" gorm:"type:decimal(7,6);"`
	Speechiness      float64   `json:"speechiness,omitempty" gorm:"type:decimal(7,6);"`
	Acousticness     float64   `json:"acousticness,omitempty" gorm:"type:decimal(7,6);"`
	Instrumentalness float64   `json:"instrumentalness,omitempty" gorm:"type:decimal(7,6);"`
	Liveness         float64   `json:"liveness,omitempty" gorm:"type:decimal(7,6);"`
	Valence          float64   `json:"valence,omitempty" gorm:"type:decimal(7,6);"`
	Tempo            float64   `json:"tempo,omitempty" gorm:"type:decimal(7,6);"`
	Duration_ms      float64   `json:"duration_ms,omitempty" gorm:"type:decimal(7,6);"`
	Time_signature   float64   `json:"time_signature,omitempty" gorm:"type:decimal(7,6);"`
}
