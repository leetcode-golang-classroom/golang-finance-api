package types

import "context"

type UserStore interface {
	GetUserByName(ctx context.Context, username string) (User, error)
	Create(ctx context.Context, username, password string) (*User, error)
}

type CategoryStore interface {
	GetAllByUserId(ctx context.Context, userId string) ([]Category, error)
	GetById(ctx context.Context, id string) (*Category, error)
	Create(ctx context.Context, userId, title string) (*Category, error)
	Delete(ctx context.Context, id string) (*Category, error)
	Update(ctx context.Context, id, title string) (*Category, error)
}
