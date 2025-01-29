package main

import (
	"database/sql"
	"fmt"
	"github.com/razorpay/movie-service/internals/config"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	_ "github.com/razorpay/movie-service/internals/database/migrations"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
		return
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	db, err := sql.Open(cfg.Database.Driver, dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <goose-command>", os.Args[0])
	}
	command := os.Args[1:]

	if err := goose.Run(command[0], db, "internals/database/migrations", command[1:]...); err != nil {
		log.Fatalf("goose run failed: %v", err)
	}
}
