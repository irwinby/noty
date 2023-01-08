package main

import (
	"context"

	"go.uber.org/zap"

	"github.com/hereisajvi/noty/internal/app"
)

// @title noty API
// @version 1.0

// @host localhost:8080
func main() {
	ctx := context.Background()

	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)

	zap.L().Info("Running application...")

	err := app.Run(ctx)
	if err != nil {
		zap.L().Fatal("failed to run app", zap.Error(err))
	}
}
