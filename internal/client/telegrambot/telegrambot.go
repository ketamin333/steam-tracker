package telegrambot

import (
	"net/http"
	"steam-tracker/internal/config"
	"time"
)

type TelegramBot struct {
	httpClient *http.Client
	BaseURL    string
	BotToken   string
}

type TelegramBotResponse struct{}

type TelegramBotCleint interface{}

var _ TelegramBotCleint = (*TelegramBot)(nil)

func New(cfg *config.Config) *TelegramBot {
	return &TelegramBot{
		httpClient: &http.Client{Timeout: time.Second * 10},
		BaseURL:    "https://api.telegram.org/bot",
		BotToken:   cfg.TelegramBotToken,
	}
}
