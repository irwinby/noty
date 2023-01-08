package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const prefix = "noty"

type (
	Server struct {
		Port string `default:"8080" split_words:"true"`
	}

	Postgres struct {
		Username                string `required:"true" split_words:"true"`
		Password                string `required:"true" split_words:"true"`
		Host                    string `default:"postgres" split_words:"true"`
		Port                    string `default:"5432" split_words:"true"`
		Database                string `required:"true" split_words:"true"`
		SSLMode                 string `default:"disable" split_words:"true"`
		MigrationsDirectoryPath string `default:"file://migrations" split_words:"true"`
	}
)

type Config struct {
	Server   *Server
	Postgres *Postgres
}

func New() (*Config, error) {
	var cfg Config

	zap.L().Info("Loading config from environment variables...")

	err := envconfig.Process(prefix, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load config from environment variables")
	}

	return &cfg, nil
}
