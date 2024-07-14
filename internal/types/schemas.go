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

type Category struct {
	Id        string     `json:"id"`
	UserId    string     `db:"user_id" json:"user_id"`
	Title     string     `json:"title"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

type Transaction struct {
	Id         string     `json:"id"`
	UserId     string     `db:"user_id" json:"user_id"`
	CategoryId string     `db:"category_id" json:"category_id"`
	Title      string     `json:"title"`
	Amount     float64    `json:"amount"`
	Currency   string     `json:"currency"`
	TxType     string     `db:"tx_type" json:"tx_type"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at"`
}
