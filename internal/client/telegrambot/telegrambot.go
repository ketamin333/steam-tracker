package telegrambot

import (
	"context"
	"net/http"
	"steam-tracker/internal/config"
	"time"
)

type TelegramBot struct {
	httpClient *http.Client
	BaseURL    string
	BotToken   string
}

type TelegramBotCleint interface {
	SendMessage(ctx context.Context, chatID int, text string) error
	SetWebhook(ctx context.Context, webhookURL string, secretToken *string) error
	DeleteWebhook(ctx context.Context) error
}

var _ TelegramBotCleint = (*TelegramBot)(nil)

func New(cfg *config.Config) *TelegramBot {
	return &TelegramBot{
		httpClient: &http.Client{Timeout: time.Second * 10},
		BaseURL:    "https://api.telegram.org/bot",
		BotToken:   cfg.TelegramBotToken,
	}
}
