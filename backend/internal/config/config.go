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
	// todo: load config from .env

	return &Config{
		Database: Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
		},
		API: API{
			Port: os.Getenv("API_PORT"),
		},
	}
}
