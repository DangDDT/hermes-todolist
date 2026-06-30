package main

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/DangDDT/hermes-todolist/backend/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: migrate [up|down|create <name>]")
		os.Exit(1)
	}

	cfg := config.MustLoad()

	m, err := migrate.New("file://migrations", cfg.Database.URL)
	if err != nil {
		fmt.Printf("failed to create migrator: %v\n", err)
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			fmt.Printf("migration up failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("migrations applied successfully")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			fmt.Printf("migration down failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("migration rolled back successfully")
	case "create":
		if len(os.Args) < 3 {
			fmt.Println("usage: migrate create <name>")
			os.Exit(1)
		}
		name := os.Args[2]
		// Create migration files manually (golang-migrate create not available in library mode).
		// We write placeholder up/down SQL files.
		upFile := fmt.Sprintf("migrations/%s.up.sql", name)
		downFile := fmt.Sprintf("migrations/%s.down.sql", name)
		if err := os.MkdirAll("migrations", 0755); err != nil {
			fmt.Printf("failed to create migrations directory: %v\n", err)
			os.Exit(1)
		}
		if err := os.WriteFile(upFile, []byte("-- migrate:up\n"), 0644); err != nil {
			fmt.Printf("failed to create %s: %v\n", upFile, err)
			os.Exit(1)
		}
		if err := os.WriteFile(downFile, []byte("-- migrate:down\n"), 0644); err != nil {
			fmt.Printf("failed to create %s: %v\n", downFile, err)
			os.Exit(1)
		}
		fmt.Printf("created migration files: %s, %s\n", upFile, downFile)
	default:
		fmt.Printf("unknown command: %s\n", command)
		os.Exit(1)
	}
}
