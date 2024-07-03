package types

import "context"

type UserStore interface {
	GetUserByName(ctx context.Context, username string) (User, error)
	Create(ctx context.Context, username, password string) (*User, error)
}
