package config

import (
	"log/slog"
	"os"
)

type Config struct {
	DBHost         string
	DBPort         string
	DBUser         string
	DBPass         string
	DBName         string
	ServerPort     string
	JWTSecret      string
	MigrationsPath string
}

func GetConfig(lgr *slog.Logger) *Config {
	lgr.Info("Getting config from .env or using default values")

	return &Config{
		DBHost:         GetOrDefault("DB_HOST", "db"),
		DBPort:         GetOrDefault("DB_PORT", "5432"),
		DBUser:         GetOrDefault("DB_USER", "postgres"),
		DBPass:         GetOrDefault("DB_PASS", "postgres"),
		DBName:         GetOrDefault("DB_NAME", "avito_shop"),
		ServerPort:     GetOrDefault("SERVER_PORT", "8080"),
		JWTSecret:      GetOrDefault("JWT_SECRET", "secret"),
		MigrationsPath: GetOrDefault("MIGRATIONS_PATH", "/app/migrations"),
	}
}

func GetOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
