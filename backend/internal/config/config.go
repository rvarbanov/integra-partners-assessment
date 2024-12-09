package config

import "os"

type Config struct {
	Database Database
	API      API
}

type Database struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type API struct {
	Port string
}

func Load() *Config {
	return &Config{
		Database: GetDBConfig(),
		API: API{
			Port: os.Getenv("API_PORT"),
		},
	}
}

func GetDBConfig() Database {
	return Database{
		Host: getEnv("DB_HOST", "localhost"),
		Port: getEnv("DB_PORT", "5432"),
		User: getEnv("DB_USER", "postgres"),
		Pass: getEnv("DB_PASSWORD", "postgres"), // Match docker-compose env var name
		Name: getEnv("DB_NAME", "userdb"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
