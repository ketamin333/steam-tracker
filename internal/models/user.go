package models

import "time"

type User struct {
	ID        int       `json:"id"`
	APIKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
}
