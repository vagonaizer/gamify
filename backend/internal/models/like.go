package models

import "time"

type Like struct {
	ID         uint64    // Unique identifier for the like
	UserID     uint64    // Receipient of the like (external key to the User table)
	TargetType string    // Type of the target (e.g., Journal, Task)
	TargetID   uint64    // ID of the target (external key to the target type)
	CreatedAt  time.Time // Timestamp of when the like was created
}
