package models

import "time"

// Journal represents a journal entry for a specific date.
type Journal struct {
	ID        uint64    // Unique identifier for the journal entry
	UserID    uint64    // External key to the User table
	Date      time.Time // Date of the journal entry
	Content   string    // Content of the journal entry
	CreatedAt time.Time // Timestamp of when the journal entry was created
	UpdatedAt time.Time // Timestamp of when the journal entry was last updated
}
