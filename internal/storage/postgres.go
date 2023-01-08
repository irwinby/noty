package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	// Must be imported to successfully apply the postgres migration.
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	// Must be imported to successfully apply the postgres migration from file source.
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// Must be imported to successfully connect to the postgres using pgx driver.
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/hereisajvi/noty/internal/config"
)

const timeout = 15 * time.Second

type Postgres struct {
	*sqlx.DB
}

func NewPostgres(ctx context.Context, cfg *config.Postgres) (*Postgres, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.Username,
		cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode)

	zap.L().Info("Establishing connection with PostgreSQL database...")

	db, err := sqlx.ConnectContext(ctx, "pgx", url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to postgres")
	}

	zap.L().Info("PostgreSQL connection established successfully")

	zap.L().Info("Creating Migrate instance to apply database migrations...")

	_migrate, err := migrate.New(cfg.MigrationsDirectoryPath, url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create migrate instance")
	}

	zap.L().Info("Applying database migrations...")

	err = _migrate.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, errors.Wrap(err, "failed to apply migrations")
	}

	zap.L().Info("Database migrations applied successfully")

	return &Postgres{
		DB: db,
	}, nil
}
