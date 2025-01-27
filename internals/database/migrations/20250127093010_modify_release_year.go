package migrations

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upModifyReleaseYear, downModifyReleaseYear)
}

func upModifyReleaseYear(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE movies
		ADD COLUMN release_year_temp VARCHAR(4);
	`)
	if err != nil {
		return fmt.Errorf("failed to add new release_year_temp column: %w", err)
	}

	_, err = tx.Exec(`
		UPDATE movies
		SET release_year_temp = release_year::VARCHAR;
	`)
	if err != nil {
		return fmt.Errorf("failed to copy data to release_year_temp: %w", err)
	}

	_, err = tx.Exec(`
		ALTER TABLE movies
		DROP COLUMN release_year;
	`)
	if err != nil {
		return fmt.Errorf("failed to drop old release_year column: %w", err)
	}

	_, err = tx.Exec(`
		ALTER TABLE movies
		RENAME COLUMN release_year_temp TO release_year;
	`)
	if err != nil {
		return fmt.Errorf("failed to rename release_year_temp to release_year: %w", err)
	}
	return nil
}

func downModifyReleaseYear(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE movies
		ADD COLUMN release_year_temp INTEGER;
	`)
	if err != nil {
		return fmt.Errorf("failed to add new release_year_temp column: %w", err)
	}

	_, err = tx.Exec(`
		UPDATE movies
		SET release_year_temp = release_year::INTEGER;
	`)
	if err != nil {
		return fmt.Errorf("failed to copy data to release_year_temp: %w", err)
	}

	_, err = tx.Exec(`
		ALTER TABLE movies
		DROP COLUMN release_year;
	`)
	if err != nil {
		return fmt.Errorf("failed to drop old release_year column: %w", err)
	}

	_, err = tx.Exec(`
		ALTER TABLE movies
		RENAME COLUMN release_year_temp TO release_year;
	`)
	if err != nil {
		return fmt.Errorf("failed to rename release_year_temp to release_year: %w", err)
	}

	return nil
}
