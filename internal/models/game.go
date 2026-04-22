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
	Name        string    `json:"name" validate:"required,max=255"`
	Description string    `json:"description" validate:"omitempty,max=4096"`
	Genre       string    `json:"genre" validate:"required,max=255"`
	Platform    Platform  `json:"platform" validate:"required,oneof=PC PlayStation Xbox Nintendo"`
	Price       float64   `json:"price" validate:"min=0"`
	Rating      float32   `json:"rating" validate:"min=0,max=10"`
	CoverURL    string    `json:"cover_url" validate:"omitempty,url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
