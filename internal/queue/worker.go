package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"steam-tracker/internal/config"

	"github.com/hibiken/asynq"
)

type Worker struct {
	srv *asynq.Server
	nt  NotifierTarget
	npc NotifierPriceChanged
}

func NewWorker(cfg *config.Config, nt NotifierTarget, npc NotifierPriceChanged) *Worker {
	return &Worker{
		srv: asynq.NewServer(
			asynq.RedisClientOpt{
				Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
				Password: cfg.RedisPassword,
			},
			asynq.Config{Concurrency: 5},
		),
		nt:  nt,
		npc: npc,
	}
}

func (w *Worker) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TypeNotificationTarget, w.handleNotificationTarget)
	mux.HandleFunc(TypeNotificationPriceChanged, w.handleNotificationPriceChanged)

	return w.srv.Start(mux)
}

func (w *Worker) Stop() {
	w.srv.Stop()
}

func (w *Worker) handleNotificationTarget(ctx context.Context, task *asynq.Task) error {
	var p PayloadNotificationTarget

	if err := json.Unmarshal(task.Payload(), &p); err != nil {
		return err
	}

	return w.nt.Send(ctx, p)
}

func (w *Worker) handleNotificationPriceChanged(ctx context.Context, task *asynq.Task) error {
	var p PayloadNotificationPriceChanged

	if err := json.Unmarshal(task.Payload(), &p); err != nil {
		return err
	}

	return w.npc.Send(ctx, p)
}
