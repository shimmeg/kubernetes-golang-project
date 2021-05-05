package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	SSL      bool
}

type Config struct {
	DB        DBConfig
	DebugMode bool
	UserRoles []string
	MaxUsers  int
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		DB: DBConfig{
			Host:     getEnv("MONGO_HOST", "localhost"),
			Port:     getEnvAsInt("MONGO_PORT", 27017),
			Username: getEnv("MONGO_USER", "aefimov"),
			Password: getEnv("MONGO_PASS", "aefimov123"),
			DBName:   getEnv("MONGO_DB", "tasks-tracker"),
			SSL:      getEnvAsBool("SSL", false),
		},
		DebugMode: getEnvAsBool("DEBUG_MODE", true),
		UserRoles: getEnvAsSlice("USER_ROLES", []string{"admin"}, ","),
		MaxUsers:  getEnvAsInt("MAX_USERS", 1),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Printf("Couldn't load env variable: %s, will use default value: %s", key, defaultVal)

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
