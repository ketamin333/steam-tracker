package pricehistoryservice

import (
	"context"
	"log/slog"
	"steam-tracker/internal/model"
	"steam-tracker/internal/queue"
	trackedgamerepo "steam-tracker/internal/repository/tracked_game"
)

func (s *Service) Run(ctx context.Context) error {
	slog.Info("run price history service")

	trackedGameRows, err := s.trackedGameRepo.GetAll(ctx)
	if err != nil {
		return err
	}

	groupedLang := make(map[string][]trackedgamerepo.TrackedGameRow)
	for _, row := range trackedGameRows {
		groupedLang[row.User.Lang] = append(groupedLang[row.User.Lang], row)
	}

	for lang, rows := range groupedLang {
		ids := make([]int, len(rows))
		gameIDs := make([]int, len(rows))
		steamRows := make(map[int]trackedgamerepo.TrackedGameRow)

		for i, row := range rows {
			ids[i] = row.Game.SteamAppID
			gameIDs[i] = row.Game.ID
			steamRows[row.Game.SteamAppID] = row
		}

		lps, err := s.priceHistoryRepo.GetLastForGames(ctx, gameIDs, lang)
		if err != nil {
			continue
		}

		lp := make(map[int]model.PriceHistory)
		for _, ph := range lps {
			lp[ph.GameID] = ph
		}

		pos, err := s.steam.AppDetails(ctx, lang, ids)
		if err != nil {
			continue
		}

		for sgID, po := range pos {
			row, ok := steamRows[sgID]
			if !ok {
				continue
			}

			ph, err := s.priceHistoryRepo.Create(
				ctx,
				&model.PriceHistory{
					GameID:          row.Game.ID,
					Lang:            row.User.Lang,
					Price:           float64(po.Final) / 100,
					Currency:        po.Currency,
					DiscountPercent: po.DiscountPercent,
				},
			)

			if err != nil {
				continue
			}

			if last, ok := lp[row.Game.ID]; ok && last.Price != ph.Price {
				if err := s.jobNotificationPriceChanged(ctx,
					queue.PayloadNotificationPriceChanged{
						User:     row.User,
						Game:     row.Game,
						OldPrice: last.Price,
						NewPrice: ph.Price,
					},
				); err != nil {
					continue
				}
			}

			if row.TargetPrice != nil && ph.Price <= *row.TargetPrice {
				if err := s.jobNotificationTarget(ctx,
					queue.PayloadNotificationTarget{
						User:         row.User,
						Game:         row.Game,
						TargetPrice:  *row.TargetPrice,
						CurrentPrice: ph.Price,
					},
				); err != nil {
					continue
				}

				if _, err := s.trackedGameRepo.Update(ctx, &row.User, row.Game.ID, nil); err != nil {
					continue
				}
			}
		}
	}

	slog.Info("ended price history service")

	return nil
}

func (s *Service) jobNotificationTarget(ctx context.Context, payload queue.PayloadNotificationTarget) error {
	return s.job.EnqueueNotificationTarget(ctx, payload)
}

func (s *Service) jobNotificationPriceChanged(ctx context.Context, payload queue.PayloadNotificationPriceChanged) error {
	return s.job.EnqueueNotificationPriceChanged(ctx, payload)
}
