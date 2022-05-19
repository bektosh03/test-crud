package postgres

import (
	"context"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	ConnString string
}

func Connect(ctx context.Context, cfg Config) (*sqlx.DB, error) {
	return sqlx.ConnectContext(ctx, "postgres", cfg.ConnString)
}
