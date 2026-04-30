package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"steam-tracker/internal/config"

	"github.com/hibiken/asynq"
)

type Worker struct {
	srv      *asynq.Server
	notifier Notifier
}

func NewWorker(cfg *config.Config, notifier Notifier) *Worker {
	return &Worker{
		srv: asynq.NewServer(
			asynq.RedisClientOpt{
				Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
				Password: cfg.RedisPassword,
			},
			asynq.Config{Concurrency: 5},
		),
		notifier: notifier,
	}
}

func (w *Worker) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TypeSendNotification, w.handleNotification)

	return w.srv.Start(mux)
}

func (w *Worker) Stop() {
	w.srv.Stop()
}

func (w *Worker) handleNotification(ctx context.Context, task *asynq.Task) error {
	var payload NotifyJobPayload

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	return w.notifier.Send(ctx, payload)
}
