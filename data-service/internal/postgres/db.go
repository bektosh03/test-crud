package postgres

import (
	"context"

	"github.com/bektosh03/test-crud/data-service/config"
	"github.com/jmoiron/sqlx"
)

func Connect(ctx context.Context, cfg config.Config) (*sqlx.DB, error) {
	return sqlx.ConnectContext(ctx, "postgres", cfg.PostgresConnString())
}
