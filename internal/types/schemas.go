package types

import (
	"database/sql"
	"time"
)

type User struct {
	Id        string       `json:"id"`
	CreatedAt time.Time    `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updated_at,omitempty"`
	Username  string       `json:"username"`
	Password  string       `json:"-"`
}
