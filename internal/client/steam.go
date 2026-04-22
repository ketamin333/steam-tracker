package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const searchURL = "https://store.steampowered.com/api/storesearch"

type SteamClient struct {
	http *http.Client
}

func NewSteamClient() *SteamClient {
	return &SteamClient{
		http: &http.Client{},
	}
}

type Price struct {
	Currency string `json:"currency"`
	Value    int    `json:"value"`
	Final    int    `json:"final"`
}

type SearchResult struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Price     *Price `json:"price"`
	TinyImage string `json:"tiny_image"`
	MetaScore string `json:"metascore"`
}

type SearchResponse struct {
	Items []SearchResult `json:"items"`
}

func (steam *SteamClient) Search(ctx context.Context, query string) ([]SearchResult, error) {
	params := url.Values{
		"term": {query},
		"l":    {"english"},
		"cc":   {"US"},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s?%s", searchURL, params.Encode()), nil)

	if err != nil {
		return nil, err
	}

	res, err := steam.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var searchResponse SearchResponse

	if err := json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		return nil, err
	}

	return searchResponse.Items, nil
}
