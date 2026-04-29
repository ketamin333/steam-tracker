package pricehistoryservice

import (
	"context"
	"log/slog"

	"steam-tracker/internal/model"
	trackedgamerepo "steam-tracker/internal/repository/tracked_game"
)

func (s *Service) Run(ctx context.Context) error {
	slog.Info("run price history service")

	tg, err := s.trackedGameRepo.GetAll(ctx)

	if err != nil {
		return err
	}

	grouped := make(map[string][]trackedgamerepo.TrackedGameRow)
	for _, row := range tg {
		grouped[row.Lang] = append(grouped[row.Lang], row)
	}

	for lang, rows := range grouped {
		ids := make([]int, len(rows))
		steamGameIDs := make(map[int]int)

		for i, row := range rows {
			ids[i] = row.SteamAppID
			steamGameIDs[row.SteamAppID] = row.GameID
		}

		po, err := s.steam.AppDetails(ctx, lang, ids)
		if err != nil {
			continue
		}

		for steamAppID, price := range po {
			gameID, ok := steamGameIDs[steamAppID]

			if !ok {
				continue
			}

			ph := model.PriceHistory{
				GameID:          gameID,
				Lang:            lang,
				Price:           float64(price.Final) / 100,
				Currency:        price.Currency,
				DiscountPercent: price.DiscountPercent,
			}

			_, err := s.priceHistoryRepo.Create(ctx, &ph)

			if err != nil {
				continue
			}
		}
	}

	slog.Info("ended price history service", "grouped", grouped)

	return nil
}
