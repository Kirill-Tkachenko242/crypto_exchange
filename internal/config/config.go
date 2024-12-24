package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort   string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	JWTSecretKey string
	APIKey       string // Ключ для доступа к внешнему API
}

func LoadConfig(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		ServerPort:   os.Getenv("SERVER_PORT"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
		APIKey:       os.Getenv("API_KEY"),
	}

	return cfg, nil
}
