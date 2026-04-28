//go:build wireinject

package app

import (
	"github.com/google/wire"
	"go.rest.api/internal/config"
	"go.rest.api/internal/db"
	"go.rest.api/internal/server"
)

func Bootstrap() (*server.Server, error) {
	wire.Build(
		config.New,
		db.New,
		repoSet,
		srvSet,
		serviceSet,
		clientSet,
		handlerSet,
	)

	return nil, nil
}
