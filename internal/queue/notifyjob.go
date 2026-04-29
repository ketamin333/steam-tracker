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

func (c *Client) Enqueue(ctx context.Context, payload NotifyJobPayload) error {
	payloadByte, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	_, err = c.client.EnqueueContext(
		ctx,
		asynq.NewTask(TypeSendNotification, payloadByte),
	)

	if err != nil {
		return err
	}

	return nil
}
