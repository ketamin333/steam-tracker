package app

import (
	"github.com/google/wire"
	userrepo "go.rest.api/internal/repository/user"
	"go.rest.api/internal/router"
	"go.rest.api/internal/server"
)

var repoSet = wire.NewSet(
	userrepo.New,
	wire.Bind(new(userrepo.UserRepository), new(*userrepo.Repository)),
)

var srvSet = wire.NewSet(
	router.New,
	server.New,
)
