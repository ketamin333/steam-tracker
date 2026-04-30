package notifier

import (
	"context"
	"steam-tracker/internal/queue"
)

type EmailTarget struct{}

func NewEmailTarget() *EmailTarget {
	return &EmailTarget{}
}

func (e *EmailTarget) Send(ctx context.Context, pay queue.PayloadNotificationTarget) error {
	return nil
}
