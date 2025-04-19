package models

import "time"

// Media represents downloaded files by the user.
type Media struct {
	ID        uint64    // Unique identifier for the media
	EntryID   *int64    // External key to the Journal table
	URL       string    // URL to the media file
	Type      string    // Type of media (e.g., image, video)
	CreatedAt time.Time // Timestamp of when the media was created
}
