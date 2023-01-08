package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/hereisajvi/noty/internal/config"
	"github.com/hereisajvi/noty/internal/delivery/http"
	v1 "github.com/hereisajvi/noty/internal/delivery/http/handler/v1"
	"github.com/hereisajvi/noty/internal/delivery/http/middleware"
	"github.com/hereisajvi/noty/internal/service"
	"github.com/hereisajvi/noty/internal/storage"
	"github.com/hereisajvi/noty/internal/storage/postgres"
)

const timeout = 5 * time.Second

func RunWithConfig(ctx context.Context, cfg *config.Config) error {
	_postgres, err := storage.NewPostgres(ctx, cfg.Postgres)
	if err != nil {
		return errors.Wrap(err, "failed to create postgres connection")
	}

	defer _postgres.Close()

	messageRepository := postgres.NewMessageRepository(_postgres)
	messageService := service.NewMessage(messageRepository)

	server := http.NewServer(
		cfg.Server,
		http.WithMiddlewares(
			middleware.
				NewLogger().
				Log(),
			middleware.
				NewRecoverer().
				Recover(),
		),
		http.WithSwagger(),
		http.WithHandlers(
			v1.NewMessageHandler(messageService),
		),
	)

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.Serve()
		if err != nil {
			zap.L().Fatal(err.Error())
		}
	}()

	<-shutdown

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to shutdown http server")
	}

	return nil
}

func Run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return errors.Wrap(err, "failed to load config")
	}

	return RunWithConfig(ctx, cfg)
}
