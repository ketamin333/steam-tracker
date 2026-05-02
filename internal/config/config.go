package config

import "os"

type Config struct {
	ServerPort       string
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	DBSSLMode        string
	RedisHost        string
	RedisPort        string
	RedisPassword    string
	MailHost         string
	MailPort         string
	MailUser         string
	MailPassword     string
	MailFrom         string
	TelegramBotToken string
}

func New() *Config {
	return &Config{
		ServerPort:       env("SERVER_PORT", "8080"),
		DBHost:           env("DB_HOST", "db"),
		DBPort:           env("DB_PORT", "5432"),
		DBUser:           env("DB_USER", ""),
		DBPassword:       env("DB_PASSWORD", ""),
		DBName:           env("DB_NAME", ""),
		DBSSLMode:        env("DB_SSL_MODE", ""),
		RedisHost:        env("REDIS_HOST", "redis"),
		RedisPort:        env("REDIS_PORT", "6379"),
		RedisPassword:    env("REDIS_PASSWORD", ""),
		MailHost:         env("MAIL_HOST", ""),
		MailPort:         env("MAIL_PORT", "465"),
		MailUser:         env("MAIL_USER", ""),
		MailPassword:     env("MAIL_PASSWORD", ""),
		MailFrom:         env("MAIL_FROM", ""),
		TelegramBotToken: env("TELEGRAM_BOT_TOKEN", ""),
	}
}

func env(key string, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
