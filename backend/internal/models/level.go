package models

type Level struct {
	ID       int64  // Unique identifier for the level
	MinPoint int    // Minimum points required to achieve this level
	MaxPoint int    // Maximum points for this level
	Name     string // Name of the level
}
