package notifier

import (
	"bytes"
	"context"
	"embed"
	"html/template"
	"steam-tracker/internal/mailer"
	"steam-tracker/internal/queue"
)

type EmailTarget struct {
	m *mailer.Mailer
}

func NewEmailTarget(m *mailer.Mailer) *EmailTarget {
	return &EmailTarget{m: m}
}

//go:embed templates/emailtarget.html
var templateEmailTarget embed.FS

func (e *EmailTarget) Send(_ context.Context, pay queue.PayloadNotificationTarget) error {
	if pay.User.Email == nil {
		return nil
	}

	tmpl, err := template.ParseFS(templateEmailTarget, "templates/emailtarget.html")
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, pay)
	if err != nil {
		return err
	}

	return e.m.Send(
		*pay.User.Email,
		"Price Alert - "+pay.Game.Name,
		buf.String(),
	)
}
