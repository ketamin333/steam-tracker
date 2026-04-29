package gameservice

import (
	"context"

	"go.rest.api/internal/client/steam"
	"go.rest.api/internal/model"
	gamerepo "go.rest.api/internal/repository/game"
)

type Service struct {
	steam steam.SteamClient
	repo  gamerepo.GameRepository
}

func New(steam steam.SteamClient, repo gamerepo.GameRepository) *Service {
	return &Service{
		steam: steam,
		repo:  repo,
	}
}

func (s *Service) Search(ctx context.Context, user *model.User, query string) ([]model.Game, error) {
	g, err := s.steam.Search(ctx, user, query)
	if err != nil {
		return nil, err
	}

	games := make([]model.Game, 0, len(g))

	for _, item := range g {
		game := model.Game{
			SteamAppID: item.ID,
			Name:       item.Name,
			CoverURL:   item.TinyImage,
		}

		if err := s.repo.Upsert(ctx, &game); err != nil {
			return nil, err
		}

		games = append(games, game)
	}

	return games, nil
}
