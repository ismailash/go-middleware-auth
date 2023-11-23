package config

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type LogFileConfig struct {
	FilePath string
}

type TokenConfig struct {
	IssuerName      string
	JwtSignatureKey []byte
	JwtLifeTime     time.Duration
}

type Config struct {
	ApiConfig
	DbConfig
	LogFileConfig
	TokenConfig
}

func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.LogFileConfig = LogFileConfig{FilePath: os.Getenv("LOG_FILE")}

	tokenLifeTime, err := strconv.Atoi(os.Getenv("TOKEN_LIFE_TIME"))
	if err != nil {
		return err
	}

	c.TokenConfig = TokenConfig{
		IssuerName:      os.Getenv("TOKEN_ISSUE_NAME"),
		JwtSignatureKey: []byte(os.Getenv("TOKEN_KEY")),
		JwtLifeTime:     time.Duration(tokenLifeTime) * time.Hour,
	}

	if c.ApiPort == "" || c.Host == "" || c.Port == "" || c.Name == "" || c.User == "" {
		return errors.New("environment required")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.readConfig(); err != nil {
		return nil, err
	}

	return cfg, nil
}
