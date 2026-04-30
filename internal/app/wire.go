//go:build wireinject

package app

import (
	"steam-tracker/internal/config"
	"steam-tracker/internal/db"
	"steam-tracker/internal/queue"
	"steam-tracker/internal/server"
	pricehistoryservice "steam-tracker/internal/service/price_history"

	"github.com/google/wire"
)

type App struct {
	Server  *server.Server
	Tracker *pricehistoryservice.Service
	Worker  *queue.Worker
}

func Bootstrap() (*App, error) {
	wire.Build(
		config.New,
		db.New,
		repoSet,
		srvSet,
		serviceSet,
		clientSet,
		handlerSet,
		queueSet,
		notifierSet,
		wire.Struct(new(App), "*"),
	)

	return nil, nil
}
