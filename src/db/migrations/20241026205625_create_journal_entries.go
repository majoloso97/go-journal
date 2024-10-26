package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateJournalEntries, downCreateJournalEntries)
}

func upCreateJournalEntries(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	create_table := `
CREATE TABLE journal_entries (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    content TEXT,
    created_at DATETIME
);`
	_, err := tx.ExecContext(ctx, create_table)
	return err
}

func downCreateJournalEntries(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	drop_table := "DROP TABLE journal_entries;"
	_, err := tx.ExecContext(ctx, drop_table)
	return err
}
