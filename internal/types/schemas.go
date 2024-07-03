package types

import "time"

type User struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
	Username  string    `json:"username"`
	Password  string    `json:"-,omitempty"`
}
