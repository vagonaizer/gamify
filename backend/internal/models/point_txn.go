// point_txn.go - Point Transaction Model
package models

import "time"

type PointTxn struct {
	ID        uint64    // Unique identifier for the point transaction
	UserID    uint64    // External key to the User table
	Amount    int       // Amount of points (positive or negative)
	CreatedAt time.Time // Timestamp of when the transaction was created
}
