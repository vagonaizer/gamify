package postgres

import (
	"context"
	"fmt"
	"gamify/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.Config) (*UserRepository, error) {
	// Constructing the DSN (Data Source Name)
	// The DSN format is: "postgres://username:password@host:port/dbname"
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	// Parsing the DSN to create a pgxpool.Config
	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %w", err)
	}

	// Creating a new connection pool
	dbpool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("unable to create db pool: %w", err)
	}

	// Testing the connection
	if err := dbpool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to connect to databse: %w", err)
	}

	// Setting the connection pool to be closed when the context is done
	return &UserRepository{
		pool: dbpool,
	}, nil
}
