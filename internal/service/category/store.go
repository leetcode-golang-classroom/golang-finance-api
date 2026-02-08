package category

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/leetcode-golang-classroom/golang-finance-api/internal/types"
)

type Store struct {
	db *sqlx.DB
}

func NewCategoryStore(db *sqlx.DB) *Store {
	return &Store{db}
}

func (store *Store) GetAllByUserId(ctx context.Context, userId string) ([]types.Category, error) {
	categories := []types.Category{}
	err := store.db.SelectContext(ctx, &categories, "SELECT * FROM categories WHERE user_id = $1", userId)
	return categories, err
}

func (store *Store) GetById(ctx context.Context, id string) (*types.Category, error) {
	category := types.Category{}
	err := store.db.SelectContext(ctx, &category, "SELECT * FROM categories WHERE id = $1", id)
	return &category, err
}

func (store *Store) Create(ctx context.Context, userId string, title string) (*types.Category, error) {
	rows, err := store.db.QueryxContext(ctx,
		`INSERT INTO categories (user_id, title)
		VALUES ($1, $2)
		RETURNING *;
		`,
		userId,
		title,
	)
	if err != nil {
		return nil, err
	}
	category := types.Category{}
	rows.Next()
	err = rows.StructScan(&category)
	return &category, err
}

func (store *Store) Delete(ctx context.Context, id string) (*types.Category, error) {
	rows, err := store.db.QueryxContext(ctx,
		`DELETE FROM categories
	 WHERE id = $1
	 RETURNING *;
	`,
		id,
	)
	if err != nil {
		return nil, err
	}
	category := types.Category{}
	rows.Next()
	err = rows.StructScan(&category)
	return &category, err
}

func (store *Store) Update(ctx context.Context, id, title string) (*types.Category, error) {
	rows, err := store.db.QueryxContext(ctx,
		`UPDATE categories
	SET title = $1
	WHERE id = $2
	RETURNING *;
	`,
		title,
		id,
	)
	if err != nil {
		return nil, err
	}
	category := types.Category{}
	rows.Next()
	err = rows.StructScan(&category)
	return &category, err
}
