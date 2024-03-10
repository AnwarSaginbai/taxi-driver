package main

import (
	"github.com/AnwarSaginbai/taxi-service/rider/internal/adapters/http"
	"github.com/AnwarSaginbai/taxi-service/rider/internal/adapters/repository/memory"
	"github.com/AnwarSaginbai/taxi-service/rider/internal/application/core/api"
	"github.com/AnwarSaginbai/taxi-service/rider/internal/config"
	"github.com/AnwarSaginbai/taxi-service/rider/internal/logger"
	"log/slog"
)

func main() {
	cfg := config.Init()

	log := logger.New()

	db := memory.New()

	app := api.NewService(db)

	adapter := http.NewAdapter(app)

	log.Info("starting the app on port", slog.Int("value", cfg.Server.Port))

	adapter.Run(cfg.Server.Port)
}
