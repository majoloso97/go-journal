package db

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	_ "majoloso97/go-journal/db/migrations"
	"majoloso97/go-journal/models"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

//go:embed migrations/*
var embedMigrations embed.FS

func GetDBConnection() *sql.DB {
	var db *sql.DB
	var err error
	db, err = sql.Open("sqlite", "journal.db")
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected")
	return db
}
func RunMigrations(db *sql.DB) {
	goose.SetDialect("sqlite")
	goose.SetBaseFS(embedMigrations)
	migrationsDir := "migrations"

	// Run migrations directly
	if err := goose.Up(db, migrationsDir); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}

func GetJournalEntries(db *sql.DB) ([]models.JournalEntry, error) {
	var entries []models.JournalEntry

	rows, err := db.Query("SELECT * FROM journal_entries")
	if err != nil {
		return nil, fmt.Errorf("GetJournalEntries: %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var entry models.JournalEntry
		if err := rows.Scan(&entry.ID, &entry.Content, &entry.Created_at); err != nil {
			return nil, fmt.Errorf("GetJournalEntries: %v", err)
		}
		entries = append(entries, entry)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetJournalentries: %v", err)
	}
	return entries, nil
}

func GetEntryByID(db *sql.DB, id int64) (models.JournalEntry, error) {
	var entry models.JournalEntry

	row := db.QueryRow("SELECT * FROM journal_entries WHERE id = ?", id)
	if err := row.Scan(&entry.ID, &entry.Content, &entry.Created_at); err != nil {
		if err == sql.ErrNoRows {
			return entry, fmt.Errorf("GetEntryByID %d: No such entry", id)
		}
		return entry, fmt.Errorf("GetEntryByID %d: %v", id, err)
	}
	return entry, nil
}

func SaveEntry(db *sql.DB, entry_str string) (models.JournalEntry, error) {
	var entry models.JournalEntry
	sql_stmt := "INSERT INTO journal_entries (content, created_at) VALUES (?, CURRENT_TIMESTAMP);"
	result, err := db.Exec(sql_stmt, entry_str)
	if err != nil {
		return entry, fmt.Errorf("SaveEntry: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return entry, fmt.Errorf("SaveEntry: %v", err)
	}
	entry, err = GetEntryByID(db, id)
	if err != nil {
		return entry, fmt.Errorf("SaveEntry: %v", err)
	}
	return entry, nil
}
