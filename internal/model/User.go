package model

import (
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID   `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v4(), not null"`
	Username string      `json:"username,omitempty" gorm:"size:255"`
	Email    string      `json:"email,omitempty" gorm:"type:varchar(100);unique_index"`
	Match    []uuid.UUID `json:"match,omitempty" gorm:"type:uuid"`
}
