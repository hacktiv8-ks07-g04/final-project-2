package config

import (
	"os"
)

type config struct {
	Database database
}

type database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Get() *config {
	return &config{
		Database: database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
	}
}
