package pricehistoryservice

import (
	"go.rest.api/internal/client/steam"
	pricehistoryrepo "go.rest.api/internal/repository/price_history"
	trackedgamerepo "go.rest.api/internal/repository/tracked_game"
)

type Service struct {
	trackedGameRepo  trackedgamerepo.TrackingRepository
	priceHistoryRepo pricehistoryrepo.PriceHistoryRepository
	steam            steam.SteamClient
}

func New(
	trackedGameRepo trackedgamerepo.TrackingRepository,
	priceHistoryRepo pricehistoryrepo.PriceHistoryRepository,
	steam steam.SteamClient,
) *Service {
	return &Service{
		trackedGameRepo:  trackedGameRepo,
		priceHistoryRepo: priceHistoryRepo,
		steam:            steam,
	}
}
