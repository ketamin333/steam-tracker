package pricehistoryservice

import (
	"steam-tracker/internal/client/steam"
	"steam-tracker/internal/queue"
	pricehistoryrepo "steam-tracker/internal/repository/price_history"
	trackedgamerepo "steam-tracker/internal/repository/tracked_game"
)

type Service struct {
	trackedGameRepo  trackedgamerepo.TrackingRepository
	priceHistoryRepo pricehistoryrepo.PriceHistoryRepository
	steam            steam.SteamClient
	job              *queue.Client
}

func New(
	trackedGameRepo trackedgamerepo.TrackingRepository,
	priceHistoryRepo pricehistoryrepo.PriceHistoryRepository,
	steam steam.SteamClient,
	job *queue.Client,
) *Service {
	return &Service{
		trackedGameRepo:  trackedGameRepo,
		priceHistoryRepo: priceHistoryRepo,
		steam:            steam,
		job:              job,
	}
}
