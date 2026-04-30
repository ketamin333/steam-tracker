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
		steamRows := make(map[int]trackedgamerepo.TrackedGameRow)

		for i, row := range rows {
			ids[i] = row.Game.SteamAppID
			steamRows[row.Game.SteamAppID] = row
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

			if row.TargetPrice != nil && ph.Price <= *row.TargetPrice {
				jp := queue.PayloadNotificationTarget{
					User:         row.User,
					Game:         row.Game,
					TargetPrice:  *row.TargetPrice,
					CurrentPrice: ph.Price,
				}

				if err := s.notifyJob.EnqueueNotificationTarget(ctx, jp); err != nil {
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
