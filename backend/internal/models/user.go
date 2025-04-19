package models

import (
	"database/sql"
	"time"
)

// User represents a user in the system.
// It contains fields for the user's ID, name, email, password, role, and date of creation
type User struct {
	ID           uint64       // Unique identifier for the user
	Email        string       // Email address of the user
	PasswordHash string       // Password for the user
	Nickname     string       // Nickname of the user
	CreatedAt    time.Time    // Timestamp of when the user was created
	UpdatedAt    time.Time    // Timestamp of when the user was last updated
	DeletedAt    sql.NullTime // nullable
}
