package model

import "time"

type User struct {
	ID        int       `json:"id"`
	APIKey    string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
