package config

import (
	"os"
)

type Config struct {
	ServerPort     string
	DatabaseURL    string
	RedisAddr      string
	RedisPass      string
	MinioEndpoint  string
	MinioAccessKey string
	MinioSecretKey string
	MinioBucket    string
	MinioUseSSL    bool
	JWTSecret      string
}

func Load() *Config {
	return &Config{
		ServerPort:     getEnv("SERVER_PORT", "8080"),
		DatabaseURL:    getEnv("DATABASE_URL", "postgres://dashboard:dashboard@postgres:5432/dashboard?sslmode=disable"),
		RedisAddr:      getEnv("REDIS_ADDR", "redis:6379"),
		RedisPass:      getEnv("REDIS_PASS", ""),
		MinioEndpoint:  getEnv("MINIO_ENDPOINT", "minio:9000"),
		MinioAccessKey: getEnv("MINIO_ACCESS_KEY", "minioadmin"),
		MinioSecretKey: getEnv("MINIO_SECRET_KEY", "minioadmin"),
		MinioBucket:    getEnv("MINIO_BUCKET", "attachments"),
		MinioUseSSL:    getEnv("MINIO_USE_SSL", "false") == "true",
		JWTSecret:      getEnv("JWT_SECRET", "change-me-in-production"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
