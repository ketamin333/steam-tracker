package model

import "time"

type PriceHistory struct {
	ID              int       `json:"id"`
	GameID          int       `json:"game_id"`
	Lang            string    `json:"lang"`
	Price           float64   `json:"price"`
	Currency        string    `json:"currency"`
	DiscountPercent int       `json:"discount_percent"`
	CheckedAt       time.Time `json:"checked_at"`
}
