package config

import (
	"os"
)


type Config struct {
	ServerAddress string
	DbHost   string
	DbPort   string
	User     string
	Password string
	DbName   string
	AccessSecret string
}

func NewConfig() (*Config, error) {
	c := Config{
		os.Getenv("SERVER_ADDRESS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("ACCESS_SECRET"),
	}
	return &c, nil
}
