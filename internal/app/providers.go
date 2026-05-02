package app

import (
	"steam-tracker/internal/client/steam"
	"steam-tracker/internal/client/telegrambot"
	gamehandler "steam-tracker/internal/handler/game"
	trackedgamehandler "steam-tracker/internal/handler/tracked_game"
	userhandler "steam-tracker/internal/handler/user"
	"steam-tracker/internal/mailer"
	"steam-tracker/internal/notifier"
	"steam-tracker/internal/queue"
	gamerepo "steam-tracker/internal/repository/game"
	pricehistoryrepo "steam-tracker/internal/repository/price_history"
	trackedgamerepo "steam-tracker/internal/repository/tracked_game"
	userrepo "steam-tracker/internal/repository/user"
	"steam-tracker/internal/router"
	"steam-tracker/internal/server"
	gameservice "steam-tracker/internal/service/game"
	pricehistoryservice "steam-tracker/internal/service/price_history"
	trackedgameservice "steam-tracker/internal/service/tracked_game"
	userservice "steam-tracker/internal/service/user"

	"github.com/google/wire"
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
	userservice.New,
	wire.Bind(new(userhandler.UserService), new(*userservice.Service)),
)

var clientSet = wire.NewSet(
	steam.New,
	wire.Bind(new(steam.SteamClient), new(*steam.Steam)),
	telegrambot.New,
	wire.Bind(new(telegrambot.TelegramBotCleint), new(*telegrambot.TelegramBot)),
)

var handlerSet = wire.NewSet(
	gamehandler.New,
	trackedgamehandler.New,
	userhandler.New,
)

var queueSet = wire.NewSet(
	queue.NewClient,
	queue.NewWorker,
)

var notifierSet = wire.NewSet(
	notifier.NewEmailTarget,
	wire.Bind(new(queue.NotifierTarget), new(*notifier.EmailTarget)),
	notifier.NewEmailPriceChanged,
	wire.Bind(new(queue.NotifierPriceChanged), new(*notifier.EmailPriceChanged)),
)

var mailerSet = wire.NewSet(
	mailer.New,
)
