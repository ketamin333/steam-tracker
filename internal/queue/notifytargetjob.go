package queue

import (
	"context"
	"encoding/json"
	"steam-tracker/internal/model"

	"github.com/hibiken/asynq"
)

const TypeNotificationTarget = "notification:target"

type PayloadNotificationTarget struct {
	User         model.User
	Game         model.Game
	TargetPrice  float64
	CurrentPrice float64
}

type NotifierTarget interface {
	Send(ctx context.Context, payload PayloadNotificationTarget) error
}

func (c *Client) EnqueueNotificationTarget(ctx context.Context, payload PayloadNotificationTarget) error {
	pb, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = c.client.EnqueueContext(ctx, asynq.NewTask(TypeNotificationTarget, pb))

	return err
}
