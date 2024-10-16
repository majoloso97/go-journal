package models

import "time"

type JournalEntry struct {
	ID         int64
	Content    string
	Created_at time.Time
}
