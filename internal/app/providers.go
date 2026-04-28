package app

import (
	"github.com/google/wire"
	"go.rest.api/internal/client"
	gamehandler "go.rest.api/internal/handler/game"
	trackedgamehandler "go.rest.api/internal/handler/tracked_game"
	gamerepo "go.rest.api/internal/repository/game"
	trackedgamerepo "go.rest.api/internal/repository/tracked_game"
	userrepo "go.rest.api/internal/repository/user"
	"go.rest.api/internal/router"
	"go.rest.api/internal/server"
	gameservice "go.rest.api/internal/service/game"
	trackedgameservice "go.rest.api/internal/service/tracked_game"
)

var repoSet = wire.NewSet(
	userrepo.New,
	wire.Bind(new(userrepo.UserRepository), new(*userrepo.Repository)),
	gamerepo.New,
	wire.Bind(new(gamerepo.GameRepository), new(*gamerepo.Repository)),
	trackedgamerepo.New,
	wire.Bind(new(trackedgamerepo.TrackingRepository), new(*trackedgamerepo.Repository)),
)

var srvSet = wire.NewSet(
	router.New,
	server.New,
)

var serviceSet = wire.NewSet(
	gameservice.New,
	wire.Bind(new(gamehandler.GameService), new(*gameservice.Service)),
	trackedgameservice.New,
	wire.Bind(new(trackedgamehandler.TrackedGameService), new(*trackedgameservice.Service)),
)

var clientSet = wire.NewSet(
	client.NewClient,
	wire.Bind(new(client.SteamSearcher), new(*client.SteamClient)),
)

var handlerSet = wire.NewSet(
	gamehandler.New,
	trackedgamehandler.New,
)
