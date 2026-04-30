package queue

import (
	"fmt"
	"steam-tracker/internal/config"

	"github.com/hibiken/asynq"
)

type Client struct {
	client *asynq.Client
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		client: asynq.NewClient(
			asynq.RedisClientOpt{
				Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
				Password: cfg.RedisPassword,
			},
		),
	}
}

func (c *Client) Close() error {
	return c.client.Close()
}
