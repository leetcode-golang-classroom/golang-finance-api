package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func Connect(uri string) (*sqlx.DB, error) {
	dataSourceName, err := pq.ParseURL(uri)
	if err != nil {
		return nil, fmt.Errorf("parse db connect string failed %w", err)
	}
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
