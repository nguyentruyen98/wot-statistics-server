package postgres

import (
	"context"
	"fmt"
	"time"
	"wot-statistics-server/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
	config config.DatabaseConfig
}

func NewDB(cfs *config.DatabaseConfig) (*DB, error) {

	if cfs == nil {
		return nil, fmt.Errorf("database config is nil")
	}

	if cfs.URL == "" {
		return nil, fmt.Errorf("database URL is empty")
	}

	db, err := sqlx.Connect("postgres", cfs.URL)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Connected to database successfully")

	return &DB{
		DB:     db,
		config: *cfs,
	}, nil

}
