//go:build wireinject

package app

import (
	"github.com/google/wire"
	"go.rest.api/internal/config"
	"go.rest.api/internal/db"
	"go.rest.api/internal/server"
	pricehistoryservice "go.rest.api/internal/service/price_history"
)

type App struct {
	Server  *server.Server
	Tracker *pricehistoryservice.Service
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
		wire.Struct(new(App), "*"),
	)

	return nil, nil
}
