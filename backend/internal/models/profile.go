package models

import "time"

// Profile broaden the User model to include additional fields for user profiles.
type Profile struct {
	UserID    uint64    // External key to the User table
	Nickname  string    // Nickname of the user
	AvatarURL string    // URL to the user's avatar image
	Bio       string    // Short biography or description of the user
	CreatedAt time.Time // Timestamp of when the profile was created
	UpdatedAt time.Time // Timestamp of when the profile was last updated
}
