package steam

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"steam-tracker/internal/model"
)

type SearchResult struct {
	Name      string  `json:"name"`
	ID        int     `json:"id"`
	TinyImage *string `json:"tiny_image"`
}

type SearchResponse struct {
	Total int            `json:"total"`
	Items []SearchResult `json:"items"`
}

func (s *Steam) Search(ctx context.Context, user *model.User, query string) ([]SearchResult, error) {
	params := url.Values{
		"term": {query},
		"l":    {"english"},
		"cc":   {user.Lang},
	}

	searchURL := fmt.Sprintf("%s/%s?%s", s.BaseURL, "storesearch", params.Encode())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, searchURL, nil)

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
