package models

import "time"

type Goal struct {
	ID          int64     // Unique identifier for the goal
	UserID      int64     // Unique identifier for the user (external key to the User table)
	Title       string    // Title of the goal
	Description string    // Description of the goal
	CreatedAt   time.Time // Timestamp of when the goal was created
}
