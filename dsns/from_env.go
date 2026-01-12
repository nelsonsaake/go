package dsns

import (
	"os"
)

func FromEnv() string {
	return New(
		WithHost(os.Getenv("DB_HOST")),
		WithPort(os.Getenv("DB_PORT")),
		WithUsername(os.Getenv("DB_USERNAME")),
		WithPassword(os.Getenv("DB_PASSWORD")),
		WithDB(os.Getenv("DB_DATABASE")),
	)
}

func ConfigFromEnv() Config {
	return Config{
		DBName:   os.Getenv("DB_DATABASE"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
	}
}
