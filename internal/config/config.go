package config

import "os"

type Config struct {
	ServerPort    string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBSSLMode     string
	RedisPassword string
}

func New() *Config {
	return &Config{
		ServerPort:    env("SERVER_PORT", "8080"),
		DBHost:        env("DB_HOST", ""),
		DBPort:        env("DB_PORT", ""),
		DBUser:        env("DB_USER", ""),
		DBPassword:    env("DB_PASSWORD", ""),
		DBName:        env("DB_NAME", ""),
		DBSSLMode:     env("DB_SSL_MODE", ""),
		RedisPassword: env("REDIS_PASSWORD", ""),
	}
}

func env(key string, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
