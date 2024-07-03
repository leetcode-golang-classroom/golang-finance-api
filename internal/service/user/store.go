package user

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/leetcode-golang-classroom/golang-finance-api/internal/types"
)

type Store struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) *Store {
	return &Store{db}
}

func (store *Store) GetUserByName(ctx context.Context, username string) (types.User, error) {
	user := types.User{}
	err := store.db.GetContext(ctx, &user, "SELECT * FROM users WHERE username = $1;", username)
	return user, err
}

func (store *Store) Create(ctx context.Context, username, password string) (*types.User, error) {
	rows, err := store.db.QueryxContext(
		ctx,
		`INSERT INTO users (username, password)
		VALUES ($1, $2)
		RETURNING *;
		`,
		username,
		password,
	)
	if err != nil {
		return nil, err
	}
	user := types.User{}
	rows.Next()
	err = rows.StructScan(&user)
	return &user, err
}
