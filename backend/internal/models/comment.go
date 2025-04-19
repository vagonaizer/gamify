package models

import "time"

type Comment struct {
	ID         uint64    // Unique identifier for the comment
	UserID     uint64    // Author of the comment (external key to the User table)
	TargetType string    // Type of the target (e.g., Journal, Task)
	TargetID   uint64    // ID of the target (external key to the target type)
	Content    string    // Content of the comment
	CreatedAt  time.Time // Timestamp of when the comment was created
}
