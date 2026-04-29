package app

import (
	"github.com/google/wire"
	"go.rest.api/internal/client/steam"
	gamehandler "go.rest.api/internal/handler/game"
	trackedgamehandler "go.rest.api/internal/handler/tracked_game"
	gamerepo "go.rest.api/internal/repository/game"
	pricehistoryrepo "go.rest.api/internal/repository/price_history"
	trackedgamerepo "go.rest.api/internal/repository/tracked_game"
	userrepo "go.rest.api/internal/repository/user"
	"go.rest.api/internal/router"
	"go.rest.api/internal/server"
	gameservice "go.rest.api/internal/service/game"
	pricehistoryservice "go.rest.api/internal/service/price_history"
	trackedgameservice "go.rest.api/internal/service/tracked_game"
)

var repoSet = wire.NewSet(
	userrepo.New,
	wire.Bind(new(userrepo.UserRepository), new(*userrepo.Repository)),
	gamerepo.New,
	wire.Bind(new(gamerepo.GameRepository), new(*gamerepo.Repository)),
	trackedgamerepo.New,
	wire.Bind(new(trackedgamerepo.TrackingRepository), new(*trackedgamerepo.Repository)),
	pricehistoryrepo.New,
	wire.Bind(new(pricehistoryrepo.PriceHistoryRepository), new(*pricehistoryrepo.Repository)),
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
	pricehistoryservice.New,
)

var clientSet = wire.NewSet(
	steam.New,
	wire.Bind(new(steam.SteamClient), new(*steam.Steam)),
)

var handlerSet = wire.NewSet(
	gamehandler.New,
	trackedgamehandler.New,
)
