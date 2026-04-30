package queue

import (
	"context"
	"encoding/json"
	"steam-tracker/internal/model"

	"github.com/hibiken/asynq"
)

const TypeNotificationPriceChanged = "notification:price_changed"

type PayloadNotificationPriceChanged struct {
	User     model.User
	Game     model.Game
	OldPrice float64
	NewPrice float64
}

type NotifierPriceChanged interface {
	Send(ctx context.Context, payload PayloadNotificationPriceChanged) error
}

func (c *Client) EnqueueNotificationPriceChanged(ctx context.Context, payload PayloadNotificationPriceChanged) error {
	pb, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = c.client.EnqueueContext(ctx, asynq.NewTask(TypeNotificationPriceChanged, pb))

	return err
}
