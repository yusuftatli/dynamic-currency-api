package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	ApplicationHost string
	ApplicationPort string
	RedisAddress    string
	Debug           bool
	BasePrefix      string
	RedisPassword   string
	RedisDb   string
	TimeInterval    string
}

// ParseEnv Get environment value from os
// If an environment required and not set raises a panic
func ParseEnv(key string, required bool, dft string) string {
	_ = godotenv.Load()
	value := os.Getenv(key)
	if value == "" && required {
		panic(fmt.Sprintf("Environment variable not found: %v", key))
	} else if value == "" {
		return dft
	}
	return value
}

func GetEnvironment() *Environment {
	return &Environment{
		RedisAddress:    ParseEnv("REDIS_ADDRESS", false, ""),
		RedisDb:    		 ParseEnv("REDIS_DB", false, "0"),
		Debug:           ParseEnv("DEBUG", false, "false") == "true",
		ApplicationHost: ParseEnv("APPLICATION_HOST", false, "0.0.0.0"),
		ApplicationPort: ParseEnv("APPLICATION_PORT", false, "8000"),
		BasePrefix:      ParseEnv("BASE_PREFIX", false, "/"),
		RedisPassword:   ParseEnv("REDIS_PASSWORD", false, ""),
		TimeInterval:    ParseEnv("TIMEINTERVAL", false, ""),
	}
}
