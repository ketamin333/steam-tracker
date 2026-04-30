package notifier

import (
	"bytes"
	"context"
	"embed"
	"html/template"
	"steam-tracker/internal/mailer"
	"steam-tracker/internal/queue"
)

type EmailPriceChanged struct {
	m *mailer.Mailer
}

func NewEmailPriceChanged(m *mailer.Mailer) *EmailPriceChanged {
	return &EmailPriceChanged{m: m}
}

//go:embed templates/emailpricechanged.html
var tplEmailPriceChanged embed.FS

func (e *EmailPriceChanged) Send(_ context.Context, pay queue.PayloadNotificationPriceChanged) error {
	if pay.User.Email == nil {
		return nil
	}

	tpl, err := template.ParseFS(tplEmailPriceChanged, "templates/emailpricechanged.html")
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, pay); err != nil {
		return err
	}

	return e.m.Send(
		*pay.User.Email,
		"Price changed - "+pay.Game.Name,
		buf.String(),
	)
}
