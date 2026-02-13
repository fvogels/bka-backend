package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)

func connectToDatabase(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", fmt.Sprintf("%s?_busy_timeout=500", path))
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Set the maximum number of idle connections to 1 to avoid issues with SQLite
	db.SetMaxOpenConns(1)

	return db, nil
}

func enableForeignKeysConstraints(db *sql.DB) error {
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return fmt.Errorf("failed to enable foreign key constraints: %w", err)
	}

	return nil
}

func setJournalMode(db *sql.DB) error {
	if _, err := db.Exec("PRAGMA journal_mode=WAL"); err != nil {
		return fmt.Errorf("failed to set journal mode to WAL: %w", err)
	}

	return nil
}

func CreateDatabase(path string) (*sql.DB, error) {
	fileExists, err := doesFileExist(path)
	if err != nil {
		return nil, fmt.Errorf("failed to create database: %w", err)
	}
	if fileExists {
		return nil, fmt.Errorf("%w: %s", ErrFileAlreadyExists, path)
	}

	database, err := connectToDatabase(path)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database stored in file %s: %w", path, err)
	}

	if err := enableForeignKeysConstraints(database); err != nil {
		return nil, err
	}

	if err := setJournalMode(database); err != nil {
		return nil, err
	}

	if err := InitializeDatabase(database); err != nil {
		return nil, err
	}

	return database, nil
}

func doesFileExist(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, fmt.Errorf("failed to check if file %s exists: %w", path, err)
}

func OpenDatabase(path string) (*sql.DB, error) {
	fileExists, err := doesFileExist(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	if !fileExists {
		return nil, fmt.Errorf("%w: %s", ErrFileDoesNotExist, path)
	}

	database, err := connectToDatabase(path)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database stored in file %s: %w", path, err)
	}

	if err := enableForeignKeysConstraints(database); err != nil {
		return nil, err
	}

	if err := setJournalMode(database); err != nil {
		return nil, err
	}

	return database, nil
}
