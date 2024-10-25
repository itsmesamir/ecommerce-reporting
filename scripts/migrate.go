package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

const (
	migrationPath = "./db/migrations"
	seedFilePath  = "./db/seeds/seeds.sql"
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func getDatabaseURL() string {
	// Get individual database configuration values
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" || password == "" || host == "" || port == "" || name == "" {
		log.Fatal("One or more required database environment variables are missing")
	}

	// Construct the database URL
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, name)
}

func runMigrationsUp() {
	cmd := exec.Command("migrate", "-path", migrationPath, "-database", getDatabaseURL(), "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to run migration up: %v", err)
	}
	log.Println("Migration up applied successfully!")

	seedDatabase()
}

func seedDatabase() {
	cmd := exec.Command("psql", getDatabaseURL(), "-f", seedFilePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}
	log.Println("Database seeded successfully!")
}

func runMigrationsDown() {
	cmd := exec.Command("migrate", "-path", migrationPath, "-database", getDatabaseURL(), "down")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to run migration down: %v", err)
	}

	log.Println("Migration down applied successfully!")
}

func main() {
	// Use os.Args to determine which migration command to run (up or down)
	if len(os.Args) < 2 {
		log.Fatal("Please specify 'up' or 'down' as the migration command.")
	}

	switch os.Args[1] {
	case "up":
		runMigrationsUp()
	case "down":
		runMigrationsDown()
	default:
		log.Fatalf("Unknown command: %s. Use 'up' or 'down'.", os.Args[1])
	}
}
