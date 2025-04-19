package models

import "time"

type Follow struct {
	FollowerID uint64    // Unique identifier for the follower
	FollowedID uint64    // Unique identifier for the followed user
	CreatedAt  time.Time // Timestamp of following
}
