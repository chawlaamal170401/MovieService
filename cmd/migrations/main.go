package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	_ "github.com/razorpay/movie-service/internals/database/migrations"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=movie_user password=password@123 dbname=movies_db sslmode=disable")
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
