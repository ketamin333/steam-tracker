package models

import "time"

type Platform string

const (
	PlatformPC          Platform = "PC"
	PlatformPlayStation Platform = "PlayStation"
	PlatformXbox        Platform = "Xbox"
	PlatformNintendo    Platform = "Nintendo"
)

type Game struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Genre       string    `json:"genre"`
	Platform    Platform  `json:"platform"`
	Price       float64   `json:"price"`
	Rating      float32   `json:"rating"`
	CoverURL    string    `json:"cover_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
