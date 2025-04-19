package models

import (
	"time"
)

// Task represent an immutable template of a task created by user
type Task struct {
	ID          uint64     // Unique identifier for the task
	UserID      int64      // External key to the User table
	Title       string     // Title of the task
	Description string     // Description of the task
	Points      int        // Points for completion of the task
	Status      TaskStatus // Status of the task (e.g., completed or not)
	CreatedAt   time.Time  // Timestamp of when the task was created
}

type TaskStatus int

const (
	TaskStatusInProgress TaskStatus = iota
	TaskStatusCompleted
	TaskStatusFailed
)
