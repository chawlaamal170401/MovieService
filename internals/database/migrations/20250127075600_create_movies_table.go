package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
	"log"
)

func init() {
	goose.AddMigrationContext(upCreateMoviesTable, downCreateMoviesTable)
}

func upCreateMoviesTable(ctx context.Context, tx *sql.Tx) error {
	query := `
        CREATE TABLE movies (
            id SERIAL PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            genre VARCHAR(100),
            director VARCHAR(255),
            release_year INT,
            rating DECIMAL(3, 1),
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `
	if _, err := tx.Exec(query); err != nil {
		log.Printf("Error creating movies table: %v", err)
		return err
	}
	return nil
}

func downCreateMoviesTable(ctx context.Context, tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS movies;`
	if _, err := tx.Exec(query); err != nil {
		log.Printf("Error dropping movies table: %v", err)
		return err
	}
	return nil
}
