package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type SteamClient struct {
	httpClient *http.Client
}

type SearchResult struct {
	Name      string  `json:"name"`
	ID        int     `json:"id"`
	TinyImage *string `json:"tiny_image"`
	Price     *Price  `json:"price"`
}

type Price struct {
	Final   int `json:"final"`
	Initial int `json:"initial"`
}

type SearchResponse struct {
	Total int            `json:"total"`
	Items []SearchResult `json:"items"`
}

type SteamSearcher interface {
	Search(ctx context.Context, query string) ([]SearchResult, error)
}

var _ SteamSearcher = (*SteamClient)(nil)

func NewClient() *SteamClient {
	return &SteamClient{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (s *SteamClient) Search(ctx context.Context, query string) ([]SearchResult, error) {
	params := url.Values{
		"term": {query},
		"l":    {"english"},
		"cc":   {"US"},
	}

	fullURL := "https://store.steampowered.com/api/storesearch?" + params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var response SearchResponse

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response.Items, nil
}
