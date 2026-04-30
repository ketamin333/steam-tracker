package mailer

import (
	"errors"
	"steam-tracker/internal/config"
	"strconv"

	"github.com/wneessen/go-mail"
)

type Mailer struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Mailer {
	return &Mailer{cfg: cfg}
}

func (m *Mailer) Send(to string, subject string, body string) error {
	if m.cfg.MailHost == "" || m.cfg.MailUser == "" || m.cfg.MailPassword == "" || m.cfg.MailFrom == "" {
		return errors.New("mail not configured")
	}

	msg := mail.NewMsg()
	msg.Subject(subject)
	msg.SetBodyString(mail.TypeTextHTML, body)

	if err := msg.From(m.cfg.MailFrom); err != nil {
		return err
	}

	if err := msg.To(to); err != nil {
		return err
	}

	port, err := strconv.Atoi(m.cfg.MailPort)

	client, err := mail.NewClient(
		m.cfg.MailHost,
		mail.WithPort(port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(m.cfg.MailUser),
		mail.WithPassword(m.cfg.MailPassword),
	)
	if err != nil {
		return err
	}

	err = client.DialAndSend(msg)

	return err
}
