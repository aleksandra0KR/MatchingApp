package user

import "github.com/gofrs/uuid/v5"

type User struct {
	ID       uuid.UUID   `json:"id,omitempty"`
	Username string      `json:"username,omitempty"`
	Email    string      `json:"email,omitempty"`
	Match    []uuid.UUID `json:"match,omitempty"`
}
