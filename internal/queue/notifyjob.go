package queue

import (
	"context"
	"encoding/json"
	"steam-tracker/internal/model"

	"github.com/hibiken/asynq"
)

const TypeSendNotification = "notification:send"

type NotifyJobPayload struct {
	User         model.User
	Game         model.Game
	TargetPrice  float64
	CurrentPrice float64
}

type Notifier interface {
	Send(ctx context.Context, payload NotifyJobPayload) error
}

func (c *Client) EnqueueNotification(ctx context.Context, payload NotifyJobPayload) error {
	payloadByte, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	_, err = c.client.EnqueueContext(
		ctx,
		asynq.NewTask(TypeSendNotification, payloadByte),
	)

	return err
}
