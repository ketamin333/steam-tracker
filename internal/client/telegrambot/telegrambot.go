package telegrambot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"steam-tracker/internal/config"
	"time"
)

type TelegramBot struct {
	httpClient *http.Client
	baseURL    string
	botToken   string
}

type TelegramBotResponse struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
	ErrorCode   int    `json:"error_code"`
	Result      any    `json:"result"`
}

type TelegramBotClient interface {
	SendMessage(ctx context.Context, chatID int64, message string, payload map[string]any) (*TelegramBotResponse, error)
}

var _ TelegramBotClient = (*TelegramBot)(nil)

func New(cfg *config.Config) *TelegramBot {
	return &TelegramBot{
		httpClient: &http.Client{Timeout: time.Second * 10},
		baseURL:    "https://api.telegram.org/bot",
		botToken:   cfg.TelegramBotToken,
	}
}

func (t *TelegramBot) doRequest(ctx context.Context, method string, payload any) (*TelegramBotResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s/%s", t.baseURL, t.botToken, method)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp TelegramBotResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, fmt.Errorf("telegram api error: %d %s", apiResp.ErrorCode, apiResp.Description)
	}

	return &apiResp, nil
}

func (t *TelegramBot) SendMessage(ctx context.Context, chatID int64, text string, options map[string]any) (*TelegramBotResponse, error) {
	payload := map[string]any{
		"chat_id": chatID,
		"text":    text,
	}

	for k, v := range options {
		payload[k] = v
	}

	return t.doRequest(ctx, "sendMessage", payload)
}
