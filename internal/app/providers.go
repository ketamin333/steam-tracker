package app

import (
	"github.com/google/wire"
	"go.rest.api/internal/client"
	gamehandler "go.rest.api/internal/handlers/game"
	gamerepository "go.rest.api/internal/repository/game"
	userrepository "go.rest.api/internal/repository/user"
)

var infrastructureSet = wire.NewSet(
	userrepository.NewUserRepository,
	gamerepository.NewGameRepository,
)

var steamSet = wire.NewSet(
	client.NewSteamClient,
)

var handlerSet = wire.NewSet(
	gamehandler.NewGameHandler,
)
