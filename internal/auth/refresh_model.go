package auth

import "time"

type RefereshToken struct {
	ID int64
	UserID int64
	Token string
	ExpiresAt time.Time
	CreatedAt time.Time
}