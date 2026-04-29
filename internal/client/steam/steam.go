package steam

import (
	"context"
	"net/http"
	"time"

	"go.rest.api/internal/model"
)

type Steam struct {
	httpClient *http.Client
	baseURL    string
}

type SteamClient interface {
	Search(ctx context.Context, user *model.User, query string) ([]SearchResult, error)
	AppDetails(ctx context.Context, lang string, steamAppID []int) (map[int]PriceOverview, error)
}

var _ SteamClient = (*Steam)(nil)

func New() *Steam {
	return &Steam{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		baseURL: "https://store.steampowered.com/api",
	}
}
