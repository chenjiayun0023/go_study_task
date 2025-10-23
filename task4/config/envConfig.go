package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

var Cfg *EnvConfig

type EnvConfig struct {
	DBPath     string
	ServerPort string
	Env        string
	JWTSecret  string
	JWTExpire  time.Duration
}

func LoadConfig() *EnvConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Warning: .env file not found:", err)
	}
	Cfg = &EnvConfig{
		Env:        getEnv("ENV", "development"),
		DBPath:     getEnv("DB_PATH", "root:root@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		JWTSecret:  getEnv("JWT_SECRET", "12345"),
		JWTExpire:  getEnvAsDuration("JWT_EXPIRE", 24*time.Hour),
	}
	fmt.Println("配置信息：", Cfg)
	return Cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	duration, err := time.ParseDuration(valueStr)
	if err != nil {
		log.Printf("Error parsing duration for %s: %v, using default", key, err)
		return defaultValue
	}
	return duration
}
