package steam

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type AppDetailsResponse struct {
	Success bool      `json:"success"`
	Data    AppDetail `json:"data"`
}

type AppDetail struct {
	PriceOverview PriceOverview `json:"price_overview"`
}

type PriceOverview struct {
	Currency         string  `json:"currency"`
	Initial          int     `json:"initial"`
	Final            int     `json:"final"`
	DiscountPercent  int     `json:"discount_percent"`
	InitialFormatted *string `json:"initial_formatted"`
	FinalFormatted   *string `json:"final_formatted"`
}

func (s *Steam) AppDetails(ctx context.Context, lang string, steamAppID []int) (map[int]PriceOverview, error) {
	ids := make([]string, len(steamAppID))
	for i, id := range steamAppID {
		ids[i] = strconv.Itoa(id)
	}

	params := url.Values{
		"appids":  {strings.Join(ids, ",")},
		"cc":      {lang},
		"filters": {"price_overview"},
	}

	appDetailsURL := fmt.Sprintf("%s/%s?%s", s.baseURL, "appdetails", params.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, appDetailsURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var raw map[string]AppDetailsResponse

	if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
		return nil, err
	}

	result := make(map[int]PriceOverview, len(raw))
	for key, val := range raw {
		if !val.Success {
			continue
		}

		id, _ := strconv.Atoi(key)
		result[id] = val.Data.PriceOverview
	}

	return result, nil
}
