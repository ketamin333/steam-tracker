//go:build wireinject

package app

import (
	"github.com/google/wire"
	"go.rest.api/internal/config"
	"go.rest.api/internal/db"
	userrepository "go.rest.api/internal/repository/user"
	"go.rest.api/internal/router"
	"go.rest.api/internal/server"
)

func Bootstrap() (*server.Server, error) {
	wire.Build(
		config.New,
		db.Init,
		userrepository.NewUserRepository,
		router.Setup,
		server.NewServer,
	)

	return nil, nil
}
