package config

import "hoc-gin/internal/utils"

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Config struct {
	DB DatabaseConfig
}

func NewConfig() *Config {
	return &Config{
		DB: DatabaseConfig{
			Host:     utils.GetEnv("DB_HOST", "localhost"),
			Port:     utils.GetEnv("DB_PORT", "5432"),
			User:     utils.GetEnv("DB_USER", "postgres"),
			Password: utils.GetEnv("DB_PASSWORD", "postgres"),
			DBName:   utils.GetEnv("DB_NAME", "myapp"),
			SSLMode:  utils.GetEnv("DB_SSLMODE", "disable"),
		},
	}
}
