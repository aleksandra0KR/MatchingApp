package model

import (
	"github.com/gofrs/uuid/v5"
)

type User struct {
	ID       uuid.UUID `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username string    `json:"username,omitempty" gorm:"size:255"`
	Email    string    `json:"email,omitempty" gorm:"type:varchar(100);unique_index"`
	Match    []User    `json:"match,omitempty" gorm:"foreignKey:id"`
}
