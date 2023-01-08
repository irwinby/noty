package http

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	swagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"

	// Must be imported to render the Swagger specification.
	_ "github.com/hereisajvi/noty/api"
	"github.com/hereisajvi/noty/internal/config"
)

const timeout = 3 * time.Second

type Handler interface {
	Routes(router chi.Router)
}

type Option interface {
	apply(router chi.Router)
}

type OptionFn func(router chi.Router)

func (fn OptionFn) apply(router chi.Router) {
	fn(router)
}

func WithSwagger() OptionFn {
	return func(router chi.Router) {
		router.Get("/swagger/*", swagger.WrapHandler)
	}
}

func WithHandlers(handlers ...Handler) OptionFn {
	return func(router chi.Router) {
		for _, handler := range handlers {
			handler.Routes(router)
		}
	}
}

func WithMiddlewares(middlewares ...func(http.Handler) http.Handler) OptionFn {
	return func(router chi.Router) {
		router.Use(middlewares...)
	}
}

type Server struct {
	server *http.Server
}

func NewServer(cfg *config.Server, opts ...Option) *Server {
	router := chi.NewRouter()

	for _, opt := range opts {
		opt.apply(router)
	}

	return &Server{
		server: &http.Server{
			Addr:              net.JoinHostPort("", cfg.Port),
			Handler:           router,
			ReadHeaderTimeout: timeout,
		},
	}
}

func (s *Server) Serve() error {
	zap.L().Info("Server is running...")

	err := s.server.ListenAndServe()
	if err != nil {
		return errors.Wrap(err, "failed to listen and serve http server")
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	zap.L().Info("Shutting down server...")

	err := s.server.Shutdown(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to shutdown http server")
	}

	return nil
}
