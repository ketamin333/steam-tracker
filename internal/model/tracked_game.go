package model

import "time"

type TrackedGame struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	GameID    int       `json:"game_id"`
	CreatedAt time.Time `json:"created_at"`
}
