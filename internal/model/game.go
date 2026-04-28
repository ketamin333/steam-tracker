package model

import "time"

type Game struct {
	ID         int       `json:"id"`
	SteamAppID int       `json:"steam_app_id"`
	Name       string    `json:"name"`
	CoverURL   *string   `json:"cover_url"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
