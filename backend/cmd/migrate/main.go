package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	// Use the same environment variables as docker-compose
	cfg := struct {
		Host   string
		Port   string
		User   string
		Pass   string
		Name   string
		InitDB string
	}{
		Host:   getEnv("DB_HOST", "localhost"), // localhost for local dev
		Port:   getEnv("DB_PORT", "5432"),
		User:   getEnv("DB_USER", "postgres"),
		Pass:   getEnv("DB_PASSWORD", "postgres"), // Match docker-compose env var name
		Name:   getEnv("DB_NAME", "userdb"),
		InitDB: "postgres", // Initial database to connect to
	}

	// First connect to postgres database to create our database
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Pass,
		cfg.InitDB,
	)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	// Create database if it doesn't exist
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", cfg.Name))
	if err != nil {
		// Ignore error if database already exists
		if !strings.Contains(err.Error(), "already exists") {
			log.Fatal(err)
		}
	}
	db.Close()

	// Now connect to our database
	dataSourceName = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Pass,
		cfg.Name,
	)

	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Get all migration files
	migrationsDir := filepath.Join("internal", "db", "migrations")
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		log.Fatal(err)
	}

	// Filter and sort migration files
	var migrations []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			migrations = append(migrations, file.Name())
		}
	}
	sort.Strings(migrations)

	// Execute migrations in order
	for _, migration := range migrations {
		fmt.Printf("Applying migration: %s\n", migration)

		migrationPath := filepath.Join(migrationsDir, migration)
		migrationSQL, err := os.ReadFile(migrationPath)
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec(string(migrationSQL))
		if err != nil {
			log.Fatalf("Error applying migration %s: %v", migration, err)
		}

		fmt.Printf("Successfully applied migration: %s\n", migration)
	}

	fmt.Println("All migrations completed successfully!")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
