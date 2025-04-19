package models

import "time"

type Notification struct {
	ID        uint64    // Unique identifier for the notification
	UserID    uint64    // External key to the User table
	Type      string    // Type of the notification (e.g., comment, like, point transaction)
	Message   string    // Message content of the notification
	IsRead    bool      // Indicates if the notification has been read
	CreatedAt time.Time // Timestamp of when the notification was created
}
