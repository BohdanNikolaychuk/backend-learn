package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresPort     string
	ServerPort       string
	JWTSecret        string
}

func LoadConfig() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		ServerPort:       os.Getenv("SERVER_PORT"),
		JWTSecret:        os.Getenv("JWT_SECRET"),
	}, nil
}
