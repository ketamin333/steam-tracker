package notifier

import (
	"context"
	"log/slog"
	"steam-tracker/internal/queue"
)

type EmailNotifier struct{}

func NewEmail() *EmailNotifier {
	return &EmailNotifier{}
}

func (e *EmailNotifier) Send(ctx context.Context, payload queue.NotifyJobPayload) error {
	slog.Info("send", "payload", payload)

	return nil
}
