package models

import "time"

type UserBadge struct {
	UserID    int64     // Unique identifier for the user (external key to the User table)
	BadgeID   int64     // Unique identifier for the badge (external key to the Badge table)
	AwardedAt time.Time // Timestamp of when the badge was awarded
}
