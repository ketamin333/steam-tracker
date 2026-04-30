package model

import "time"

type User struct {
	ID        int       `json:"id"`
	APIKey    string    `json:"-"`
	Lang      string    `json:"lang"`
	Email     *string   `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
