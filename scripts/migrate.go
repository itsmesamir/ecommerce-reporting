package main

import (
	"log"
	"os"
	"os/exec"
)

const (
	migrationPath = "./db/migrations"
	seedFilePath  = "./db/seeds/seeds.sql"
	databaseURL   = "postgres://samiralam:password@localhost:5432/ecommerce_reporting?sslmode=disable"
)

func runMigrationsUp() {
	cmd := exec.Command("migrate", "-path", migrationPath, "-database", databaseURL, "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to run migration up: %v", err)
	}
	log.Println("Migration up applied successfully!")

	seedDatabase()
}

func seedDatabase() {
	cmd := exec.Command("psql", databaseURL, "-f", seedFilePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}
	log.Println("Database seeded successfully!")
}

func runMigrationsDown() {
	cmd := exec.Command("migrate", "-path", migrationPath, "-database", databaseURL, "down")
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
