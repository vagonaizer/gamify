package models

import "time"

type Badge struct {
	ID          int64     // Unique identifier for the badge
	Name        string    // Name of the badge
	Description string    // Description of the badge
	CreatedAt   time.Time // Timestamp of when the badge was created
}
