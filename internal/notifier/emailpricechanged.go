package notifier

import (
	"context"
	"steam-tracker/internal/queue"
)

type EmailPriceChanged struct{}

func NewEmailPriceChanged() *EmailPriceChanged {
	return &EmailPriceChanged{}
}

func (e *EmailPriceChanged) Send(ctx context.Context, pay queue.PayloadNotificationPriceChanged) error {
	return nil
}
